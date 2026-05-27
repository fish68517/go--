<template>
  <div>
    <div v-if="showForm" id="btda">
      <div class="row">
        <div class="section-title col-sm-8 col-sm-offset-2">
          <h2 class="text-uppercase branding-color" style="text-align:center;padding-top:20px">
            基于联盟链fabrci的智慧中药房-汤药微信端查询
            <span>
              <i class="fa fa-circle"></i>
              <i class="fa fa-circle"></i>
              <i class="fa fa-circle"></i>
            </span>
          </h2>
          <p style="text-align:center;padding-bottom:20px">智慧中药房提供最近半年上的汤药记录查询</p>
          <el-card>
            <el-form :inline="true" :model="queryParam" ref="queryParam" class="demo-form-inline">
              <el-form-item label="医院处方号：" prop="recipe_code">
                <el-input v-model="queryParam.recipe_code" placeholder="医院处方号"  style="width:400px"></el-input>
              </el-form-item>
            </el-form>
            <div style="text-align: center">
              <button @click="fetchMedicineProcess" class="query-button">查询</button>
            </div>
          </el-card>
        </div>
      </div>
    </div>
    <div v-else>
      <div v-if="hasProcesses">
        <!-- 显示查询结果 -->
        <div class="timeline">
          <div v-for="(process, index) in processes" :key="index" class="timeline-item">
            <div class="timeline-circle"></div>
            <div class="timeline-content">
              <h3>{{ process.Step }}</h3>
              <p>{{ process.Performer }}</p>
              <span>{{ process.Timestamp }}</span>
            </div>
          </div>
        </div>
        <section v-if="medicines.length" class="trace-panel">
          <h3>药材来源明细</h3>
          <table class="trace-table">
            <thead>
              <tr>
                <th>药材</th>
                <th>编码</th>
                <th>源地</th>
                <th>批次</th>
                <th>数量</th>
                <th>单价</th>
                <th>小计</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="(item, index) in medicines" :key="(item.drug_product_number || '') + '-' + (item.batch_no || '') + '-' + index">
                <td>{{ item.drug_product_name }}</td>
                <td>{{ item.drug_product_number }}</td>
                <td>{{ item.purchase_origin }}</td>
                <td>{{ item.batch_no }}</td>
                <td>{{ item.quantity }}</td>
                <td>{{ item.unit_price }}</td>
                <td>{{ item.total_price }}</td>
              </tr>
            </tbody>
          </table>
        </section>
        <section v-if="payment.pay_status || blockchain.chain_status" class="trace-panel">
          <h3>支付与上链信息</h3>
          <div class="trace-meta">
            <p>支付状态：{{ payment.pay_status || '未记录' }}</p>
            <p>应付金额：{{ payment.total_amount || 0 }} 元</p>
            <p>实付金额：{{ payment.pay_amount || 0 }} 元</p>
            <p>支付方式：{{ payment.pay_method || '未记录' }}</p>
            <p>模拟交易号：{{ payment.mock_trade_no || '未记录' }}</p>
            <p>支付时间：{{ payment.paid_at || '未记录' }}</p>
            <p>上链状态：{{ blockchain.chain_status || '未上链' }}</p>
            <p>交易ID：{{ blockchain.tx_id || '未记录' }}</p>
            <p>Payload Hash：{{ blockchain.payload_hash || '未记录' }}</p>
          </div>
        </section>
      </div>
      <div v-else>
        <!-- 无查询结果时显示提示信息 -->
        <p style="text-align:center; color:red;">无该病人信息, 请检查信息是否正确!</p>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

// 配置axios默认基础路径（配合代理使用）
axios.defaults.baseURL = '/api';

export default {
  name: 'App',
  data() {
    return {
      showForm: true,
      queryParam: {
        recipe_code: '',
      },
      processes: [],
      medicines: [],
      payment: {},
      blockchain: {},
    };
  },
  computed: {
    hasProcesses() {
      return this.processes.length > 0;
    }
  },
  methods: {
    // 新增：数据转换函数 - 将原始数据转为Timeline需要的格式
    transformProcessData(rawData) {
      if (Array.isArray(rawData?.process)) {
        return rawData.process.map(item => ({
          Step: item.Step || item.step,
          Performer: item.Performer || item.performer || '未记录',
          Timestamp: this.formatDisplayTime(item.Timestamp || item.timestamp),
        })).filter(item => item.Timestamp);
      }
      // 步骤映射配置：按【接方、审核、泡药、煎药、包装、发货】的固定顺序
      const stepConfig = [
        { stepName: '接方', performerKey: 'do_person', timeKey: 'do_time' },
        { stepName: '审核', performerKey: 'pres_aduit_reviewer', timeKey: 'pres_aduit_time' },
        { stepName: '泡药', performerKey: 'soaking_person', timeKey: 'soak_start_time' },
        { stepName: '煎药', performerKey: 'decoction_person', timeKey: 'decoction_start_time' },
        { stepName: '包装', performerKey: 'packing_personnel', timeKey: 'pack_start_time' },
        { stepName: '发货', performerKey: 'delivery_personnel', timeKey: 'delivery_time' },
      ];

      // 处理时间格式：去除秒级，保留 "YYYY-MM-DD HH:MM"
      // 生成目标格式数组
      return stepConfig.map(item => {
        // 处理空值：执行人为空时显示"未记录"
        const performer = rawData[item.performerKey] || '未记录';
        const timestamp = this.formatDisplayTime(rawData[item.timeKey]);
        
        return {
          Step: item.stepName,
          Performer: performer,
          Timestamp: timestamp
        };
      }).filter(item => item.Timestamp);
    },

    formatDisplayTime(timeStr) {
      if (!timeStr) return '';
      const cleanTime = String(timeStr).replace('T', ' ').replace(/\s+/g, ' ').trim();
      return cleanTime.split(':').slice(0, 2).join(':').replace(' ', ' ');
    },

    async fetchMedicineProcess() {
      // 1. 强制转换为字符串并去空格，确保参数纯净
      const recipeCode = this.queryParam.recipe_code.trim();
      console.log("cfh==", recipeCode);
      
      // 2. 更严格的参数校验
      if (!recipeCode) {
        this.$message?.error('请输入医院处方号') || alert('请输入医院处方号');
        return;
      }

      try {
        // 3. 使用代理路径，不再直接拼接后端IP
        const response = await axios.get('/queryRecipes', {
          params: {
            recipe_code: recipeCode
          }
        });
        
        console.log(response.data?.code);
        if (response.status === 200 && (response.data?.code === 200 || response.data?.code === 0)) {
          // 核心修改：调用转换函数处理数据
          let data = response.data?.data || {};
          if (typeof data === 'string') {
            try {
              data = JSON.parse(data);
            } catch (e) {
              data = {};
            }
          }
          if (Array.isArray(data)) {
            data = data[0] || {};
          }
          this.processes = this.transformProcessData(data);
          this.medicines = data.medicines || [];
          this.payment = data.payment || {};
          this.blockchain = data.blockchain || {};
          this.showForm = false;
        } else {
          console.error('查询失败:', response.data);
          this.$message?.error('查询失败：' + (response.data?.msg || '服务器返回异常')) || alert('查询失败，请稍后再查询');
        }

      } catch (error) {
        console.error('请求异常:', error);
        // 分场景提示错误
        if (error.response?.status === 500) {
          this.$message?.error('服务器内部错误，请联系管理员') || alert('服务器内部错误，请联系管理员');
        } else if (error.message.includes('Network')) {
          this.$message?.error('网络异常，请检查网络连接') || alert('网络异常，请检查网络连接');
        } else {
          this.$message?.error('查询失败，请检查处方号是否正确') || alert('查询失败，请检查处方号是否正确');
        }
      }
    },
  },
};
</script>

<style>
[v-cloak] {
  display: none;
}
.td1 {
  text-align: center;
  width: 6%;
}
.td2 {
  text-align: left;
  width: 8%;
}
.td3 {
  text-align: center;
  width: 10%;
}
#app {
  padding: 20px;
  font-family: Avenir, Helvetica, Arial, sans-serif;
  text-align: center;
}
.search-form {
  margin-bottom: 20px;
}
.query-button {
  padding: 10px 20px;
  background-color: #42b983;
  color: white;
  border: none;
  cursor: pointer;
  border-radius: 4px; /* 增加圆角，提升样式 */
}
.query-button:hover {
  background-color: #359469; /* 增加hover效果 */
}
.timeline {
  position: relative;
  max-width: 1080px;
  margin: 24px auto;
  padding-left: 40px;
  text-align: left;
}
.timeline::before {
  content: '';
  position: absolute;
  top: 0;
  left: 24px;
  height: 100%;
  width: 6px;
  background: linear-gradient(to bottom, #34a853, #005f4f);
  z-index: -1;
}
.timeline-item {
  position: relative;
  margin-bottom: 42px;
}
.timeline-circle {
  position: absolute;
  left: 0;
  top: 20px;
  width: 48px;
  height: 48px;
  background-color: #34a853;
  border: 4px solid white;
  border-radius: 50%;
  box-shadow: 0 0 0 6px rgba(52, 168, 83, 0.3);
}
.timeline-content {
  padding-left: 70px;
}
.timeline-content h3 {
  margin: 0;
  font-size: 1.75em;
  font-weight: bold;
  color: #333;
}
.timeline-content p {
  margin: 10px 0;
  color: #666;
}
.timeline-content span {
  color: #999;
  font-size: 0.95em;
}
.trace-panel {
  max-width: 1080px;
  margin: 24px auto;
  text-align: left;
}
.trace-panel h3 {
  margin: 0 0 12px;
  color: #23443a;
}
.trace-table {
  width: 100%;
  border-collapse: collapse;
  background: #fff;
}
.trace-table th,
.trace-table td {
  border: 1px solid #d9e5df;
  padding: 10px 12px;
  text-align: center;
}
.trace-table th {
  background: #eff7f2;
  color: #23443a;
}
.trace-meta {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(260px, 1fr));
  gap: 10px 18px;
  padding: 16px;
  border: 1px solid #d9e5df;
  background: #fff;
}
.trace-meta p {
  margin: 0;
  color: #2f3f39;
  word-break: break-all;
}
</style>
