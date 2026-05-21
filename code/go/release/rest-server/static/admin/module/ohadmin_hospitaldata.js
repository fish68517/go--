
layui.use(['func'], function () {

    //声明变量
    var func = layui.func
        ,form = layui.form
        , $ = layui.$;

    if (A == 'index') {
        //【TABLE列数组】
        var cols = [
            {type: 'checkbox', fixed: 'left'}
            , {field: 'hospital_name', width: 250, title: '医院名称', align: 'center'}
            , {field: 'prescription_count', width: 100, title: '开方数量', align: 'center'}
            , {field: 'prescription_fee', width: 250, title: '加工费用', align: 'center'}
            , {field: 'drug_fee', width: 100, title: '药品费用', align: 'center'}
            , {field: 'prescription_total_fee', width: 100, title: '处方总金额', align: 'center'}

        ];

        //【渲染TABLE】
        func.tableIns(cols, "tableList");

        //【设置弹框】
        func.setWin("医院数据汇总");

    }
});
