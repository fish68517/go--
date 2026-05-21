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
        <Timeline :processes="processes" />
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
import Timeline from './components/Timeline.vue';

// 配置axios默认基础路径（配合代理使用）
axios.defaults.baseURL = '/api';

export default {
  name: 'App',
  components: {
    Timeline,
  },
  data() {
    return {
      showForm: true,
      queryParam: {
        recipe_code: '',
      },
      processes: [],
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
      const formatTime = (timeStr) => {
        if (!timeStr) return '';
        // 先去除多余空格，再截取到分钟级
        const cleanTime = timeStr.replace(/\s+/g, ' ').trim();
        return cleanTime.split(':').slice(0, 2).join(':').replace(' ', ' ');
      };

      // 生成目标格式数组
      return stepConfig.map(item => {
        // 处理空值：执行人为空时显示"未记录"
        const performer = rawData[item.performerKey] || '未记录';
        const timestamp = formatTime(rawData[item.timeKey]);
        
        return {
          Step: item.stepName,
          Performer: performer,
          Timestamp: timestamp
        };
      });
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
        if (response.status === 200 && response.data?.code === 200) {
          // 核心修改：调用转换函数处理数据
          this.processes = this.transformProcessData(response.data?.data || {});
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
</style>