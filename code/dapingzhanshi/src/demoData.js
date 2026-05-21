export default {
    //市场趋势数据
    marketData: [{
        name: '处方贴数',
        data: [{
            title: '1月',
            value: 438
        }, {
            title: '2月',
            value: 1127
        }, {
            title: '3月',
            value: 3431
        }, {
            title: '4月',
            value: 649.8
        }]
    }],
    //雷管使用数据
    useData: [{
        name: '调剂类型占比',
        data: [{
            title: '代煎',
            value: 424
        }, {
            title: '代配',
            value: 256
        }, {
            title: '膏方',
            value: 109
        }]
    }],
    //客户销售数据
    customerData: [{
        name: '2020客户销售数据',
        data: [ {
            title: '西安庆华',
            value: 90
        }]
    }],
    //生产监控数据
    produceDaata: [{
        name: 'top10 饮品使用量 ',
        data: [{
            title: '2020/01',
            value: 300
        }, {
            title: '2020/02',
            value: 300
        }, {
            title: '2020/03',
            value: 400
        }, {
            title: '2020/04',
            value: 500
        }, {
            title: '2020/05',
            value: 700
        }]
    }],
    //2020销售情况数据
    saleData: [{
        name: '计划销售量',
        data: [{
            title: '2020/01',
            value: 175
        }, {
            title: '2020/02',
            value: 135
        }, {
            title: '2020/03',
            value: 295
        }, {
            title: '2020/04',
            value: 375
        }, {
            title: '2020/05',
            value: 520
        }]
    }],
    //整体概览
    overviewData: [{
        name: '整体概览',
        data: [{
            title: '爆破企业',
            value: 936
        }, {
            title: '爆破员',
            value: 10947
        }, {
            title: '起爆器',
            value: 10414
        }, {
            title: '作业点',
            value: 37923
        }, {
            title: 'APP装机量',
            value: 11489
        }]
    }],
    //地图展示数据
    mapData: [{
        data: [{
            "hospital_name": "河南省人民医院",
            "decoction_type": "代煎",
            "prescription_number": "20780",
            "patient_name": "静态数据",
            "dosage": 7,
            "current_state": "发货",
            "drug_count": 0
        }]
    }],
    totalData: [{
        name: '大屏总览数据',
        produceTotal: 65770000,//生产总量
        useTotal: 61403314,//使用总量
        modalSale: 6003,//模块销售额
        detonatorSale: 130,//起爆器销售额
        equipmentSale: 136,//生产仪器设备销售额
        materialSale: 535//辅材销售额
    }]
}
