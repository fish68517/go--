import Vue from 'vue'
import Vuex from 'vuex'
import { getBigScreenData } from '../api/api';

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    //市场趋势数据
    marketData: [{
      name: '处方贴数',
      data: [{
        title: '2025-1-1',
        value: 1
      }, {
        title: '2025-1-2',
        value: 2
      }, {
        title: '2025-1-3',
        value: 6
      }, {
        title: '2025-1-4',
        value: 4
      }, {
        title: '2025-1-5',
        value: 3
      }, {
        title: '2025-1-6',
        value: 12
      }, {
        title: '2025-1-7',
        value: 20
      }]
    }],
    //雷管使用数据
    useData: [{
      name: '调剂类别统计',
      data: [{
        title: '代煎',
        value: 30
      }, {
        title: '代配',
        value: 70
      }]
    }],
    //客户销售数据
    customerData: [{
      name: '医院处方占比',
      data: [ {
        title: '河南省人民医院 ',
        value: 0
      }, {
        title: '郑州人民医院',
        value: 0
      }, {
        title: '河南省中医院附属医院',
        value: 0
      }]
    }],
    //生产监控数据
    produceData: [{
      name: 'top10 饮品使用量 ',
      data: [ {
        title: '2021/02',
        value: 0
      }, {
        title: '2021/03',
        value: 0
      }, {
        title: '2021/04',
        value: 0
      }, {
        title: '2021/05',
        value: 0
      }, {
        title: '2021/06',
        value: 0
      }, {
        title: '2021/06',
        value: 0
      }, {
        title: '2021/07',
        value: 0
      }, {
        title: '2021/08',
        value: 0
      }, {
        title: '2021/09',
        value: 0
      }, {
        title: '2021/10',
        value: 0
      }]
    }],
    //2021销售情况数据
    saleData: [{
      name: '接方',
      data: [{
        title: '2021/02',
        value: 0
      }, {
        title: '2021/03',
        value: 0
      }, {
        title: '2021/04',
        value: 0
      }, {
        title: '2021/05',
        value: 0
      }, {
        title: '2021/06',
        value: 0
      }]
    }, {
      name: '完成',
      data: [{
        title: '2021/02',
        value: 0
      }, {
        title: '2021/03',
        value: 0
      }, {
        title: '2021/04',
        value: 0
      }, {
        title: '2021/05',
        value: 0
      }, {
        title: '2021/06',
        value: 0
      }]
    }],
    //整体概览
    overviewData: [{
      name: '整体概览',
      data: [{
        title: '爆破企业',
        value: 0
      }, {
        title: '爆破员',
        value: 0
      }, {
        title: '起爆器',
        value: 0
      }, {
        title: '作业点',
        value: 0
      }, {
        title: 'APP装机量',
        value: 0
      }]
    }],
    //地图展示数据
    mapData: [{
      name: '处方数据',
      data: [{
          "hospital_name": "河南省人民医院",
          "decoction_type": "代煎",
          "prescription_number": "20780",
          "patient_name": "静态数据",
          "dosage": 7,
          "current_state": "发货",
          "drug_count": 0
        }],
    }],
    //总体数据
    totalData: [{
      name: '大屏总览数据',
      produceTotal: 0,//生产总量
      useTotal: 0,//使用总量
      modalSale: 0,//模块销售额
      detonatorSale: 0,//起爆器销售额
      equipmentSale: 0,//生产仪器设备销售额
      materialSale: 0//辅材销售额
    }],
    //ajax获取的大屏数据
    ajaxData:null
  },
  mutations: {
    changeMarketData(state, marketData) {
      state.marketData = marketData;
    },
    changeUseData(state, useData) {
      state.useData = useData;
    },
    changeCustomerData(state, customerData) {
      state.customerData = customerData;
    },
    changeProduceData(state, produceData) {
      state.produceData = produceData;
    },
    changeSaleData(state, saleData) {
      state.saleData = saleData;
    },
    changeOverviewData(state, overviewData) {

      state.overviewData = overviewData;
    },
    changeMapData(state, mapData) {
 console.log("changeMapData", JSON.stringify(mapData));
      state.mapData = mapData;
    },
    changeTotalData(state, totalData) {
      console.log("changeTotalData", JSON.stringify(totalData));
      state.totalData = totalData;
    },
    changeAjaxData(state,ajaxData){
      state.ajaxData = ajaxData;
    }
  },
  actions: {
    //异步设置大屏数据
    actionBigScreenData(context){
			let isdemo = window._DEMODATA['isDemodata'];
			if(isdemo){//如果使用demo假数据就不请求后台数据。
				const demodata = window._DEMODATA['bigScreenDataDemo'];
				context.commit('changeAjaxData', demodata);
				return
			}
      getBigScreenData().then(res => {
        if(res.code===200){
          let stateData = res.result.state;

          context.commit('changeAjaxData', stateData);
        }
      }).catch(err => {
        console.log(err.type);
      });
    },

    //异步设置市场趋势数据
    actionMarketData(context,params) {
      let datas = params?params:context.state.marketData;
      context.commit('changeMarketData', datas);
    },

    //异步设置雷管使用量排行数据
    actionUseData(context,params) {
      let datas = params?params:context.state.useData;
      context.commit('changeUseData', datas);
    },

    actionCustomerData(context,params) {
      let arr = params?params:context.state.customerData;
      context.commit('changeCustomerData', arr);
    },


    actionProduceData(context,params) {
      let arr = params?params:context.state.produceData;
      context.commit('changeProduceData', arr);
    },


    actionSaleData(context,params) {
      let arr = params?params:context.state.saleData;
      console.log("asyc-actionSaleData --data");
      context.commit('changeSaleData', arr);
    },

    //异步设置整体概览数据
    actionOverviewData(context,params) {
      let datas = context.state.overviewData;
      if(params){
        datas = [{
          name: '整体概览',
          data: params
        }];
      }
      console.log("action-actionOverviewData --data", datas);
      context.commit('changeOverviewData', datas);
    },
    //异步设置地图展示数据
    actionMapData(context,params) {
      let datas = params?[params]:context.state.mapData;
      // console.log("actionMapData"+arr);
      // context.commit('changeMapData', arr);
      if(params){
        datas = [{
          name: '处方列表',
          data: params
        }];
      }
      console.log("action actionMapData  --data"+JSON.stringify(datas));
      context.commit('changeMapData', datas);
    },

    //异步设置总览数据
    actionTotalData(context,params) {
      let arr = params?[params]:context.state.totalData;
      console.log("action actionTotalData  --data"+JSON.stringify(arr));
      context.commit('changeTotalData', arr);
    }
  },
  getters:{
    overviewData: state=>{
      let datas = [];
      if(state.ajaxData && state.ajaxData.overviewData){
        datas = state.ajaxData.overviewData;
      }
      return {
        name: '整体概览',
        data: datas
      }
    },
    mapData: state=>{
      let datas = [];
      if(state.ajaxData && state.ajaxData.mapData){
        datas = state.ajaxData.mapData;
      }
      return {
        name: '处方列表',
        data: datas
      }
    }
  }
})
