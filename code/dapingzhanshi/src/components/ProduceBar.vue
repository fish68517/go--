<template>
  <Echarts
    class="myBar"
    title="top10 饮品使用量 (单位：kg)"
    classify="bar"
    :dataArray="dataArr"
    :isBorder=false
    :isHorizontal=true
    :colors='colors'
    :isLegend=true
  ></Echarts>
</template>

<script>
import Echarts from "../echarts/Echarts";
export default {
  name: "ProduceBar",
  components: {
    Echarts
  },
  data() {
    return {
      colors:['#65f5f3','#faad14']
    };
  },
  computed:{
    dataArr(){
      return this.$store.state.produceData;
    }
  }
};
</script>
<style lang="less" scoped>
.myBar {
  flex: 1;
  background-color: rgba(41, 52, 65, 0.2);
  border-radius: 5px;
}
</style>
