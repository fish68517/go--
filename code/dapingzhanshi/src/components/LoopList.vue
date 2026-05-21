<template>
  <div class="wrap">
    <ul class="title">
      <li>来源机构</li>
      <li>处方号</li>
      <li>患者姓名</li>
      <li>调配类型</li>
      <li>贴数</li>
      <li>药品味数</li>
      <li>流程环节</li>
    </ul>
    <ul class="list">
      <li class="list-item" :class="{ 'hover': isAnimation }" v-for="item in showData" :key="item.prescription_number">
        <div>{{item.hospital_name}}</div>
        <div>{{item.prescription_number}}</div>
        <div>{{desensitizeName(item.patient_name)}}</div>
        <div>{{item.decoction_type}}</div>
        <div>{{item.dosage}}</div>
        <div>{{item.drug_count}}</div>
        <div>{{item.current_state}}</div>
      </li>
    </ul>
  </div>
</template>

<script>


export default {
  name: "LoopList",
  data() {
    return {
      isAnimation: false,
      timeout: null,
      timeoutA: null,
      pageIndex: 1,
      showData: []
    };
  },
  computed:{
    dataArr(){
      console.log("looplist-computed"+this.timeString);
      return this.$store.state.mapData
    }
  },
  mounted() {
    this.goLoop();
  },
  beforeDestroy(){
   console.log("looplist-beforeDestroy"+this.timeString);
    clearInterval(this.timeout);
    clearTimeout(this.timeoutA);
  },
  methods: {
desensitizeName(name){
  if (name.length<=2){
    return name.split("").map(()=>'*').join("");
  }
  return name.charAt(0)+"*".repeat(name.length-2)+name.charAt(name.length-1);
},
    //启用滚屏
    goLoop(){
      this.showData = this.dataArr.slice(0, 14);
      this.timeout = setInterval(() => {
        this.animationFn();
        clearTimeout(this.timeoutA);
        this.timeoutA = setTimeout(()=>{
          this.paging();
        },1000);
      }, 5000);
    },
    getTime() {
      let time = new Date();
      let y = time.getFullYear();
      let m = time.getMonth() + 1;
      m = m > 9 ? m : "0" + m;
      let d = time.getDate();
      d = d > 9 ? d : "0" + d;
      let hh = time.getHours();
      hh = hh > 9 ? hh : "0" + hh;
      let mm = time.getMinutes();
      mm = mm > 9 ? mm : "0" + mm;
      let ss = time.getSeconds();
      ss = ss > 9 ? ss : "0" + ss;
      let str = `${y}年${m}月${d}日 ${hh}:${mm}:${ss}`;
      this.timeString = str;
    },
    //滚屏分页
    paging() {
      this.pageIndex++;
      let arr = this.dataArr[0].data;
      let endNum = this.pageIndex * 14;
      let startNum = (this.pageIndex-1)*14
      this.showData = arr.slice(startNum, endNum);
      if (this.showData.length < 14) {
        this.pageIndex = 0;
      }

    },
    //启用动画
    animationFn() {
      if (this.isAnimation) {
        this.isAnimation = false;
      } else {
        this.isAnimation = true;
      }
    }
  }
};
</script>
<style lang="less" scoped>
.wrap {
  flex: 1;
  background: rgba(0, 0, 0, 0.1);
  border-radius: 8px;
  display: flex;
  flex-direction: column;
  padding: 0 5px;
  border-bottom: 2px solid #9cefff;
  border-top: 1px solid #ffd792;
  box-shadow: 0 0 1px 0 #f5f5f5;
  .title {
    width: 100%;
    color: rgba(255, 255, 255, 0.8);
    text-shadow: 1px 1px 0 #e4393c;
    display: flex;
    flex-direction: row;
    padding: 5px 0;
    border-bottom: 1px solid #aaaaaa;
    li {
      flex: 1;
      text-align: center;
      font-size: 12px;
      font-weight: bold;
    }
  }
  .list {
    flex: 1;
    display: flex;
    flex-direction: column;
    .list-item {
      width: 100%;
      margin-top: 5px;
      display: flex;
      flex-direction: row;
      background-color: rgba(255, 255, 255, 0.1);
      border-radius: 5px;
      padding: 1px 0;
      font-size: 12px;
      color: #ffffff;
      text-shadow: 1px 1px 1px #11bcff;
      transform: rotateX(0deg);
      transition: transform 2s;
      div {
        flex: 1;
        text-align: center;
      }
    }
    .hover {
      transform: rotateX(360deg);
    }
  }
}
</style>
