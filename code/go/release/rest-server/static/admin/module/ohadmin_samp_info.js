/**
 * 泡药管理
 * @author fengxh
 * @since 2025/1/13
 */
layui.use(['func'], function () {

    //声明变量
    var func = layui.func
        ,form = layui.form
        , $ = layui.$;

    if (A == 'index') {
        //【TABLE列数组】
        var cols = [
            {type: 'checkbox', fixed: 'left'}
            , {field: 'prescription_number', width: 180, title: '处方号', align: 'center'}
            , {field: 'prescription_weight', width: 90, title: '处方重量', align: 'center'}
            , {field: 'prescription_sj_weight', width: 90, title: '实际重量', align: 'center'}
            , {field: 'drug_count', width: 100, title: '药品数量', align: 'center'}
            , {field: 'drug_sj_count', width: 120, title: '药品实际数量', align: 'center'}
            , {field: 'dosage', width: 100, title: '贴数', align: 'center'}
            , {field: 'remark', width: 100, title: '检测人员', align: 'center'}
            , {field: 'operate_name', width: 100, title: '检测人员', align: 'center'}
            , {field: 'operate_time', width: 180, title: '检测时间', align: 'center', templet:"<div>{{layui.util.toDateString(d.operate_time*1000, 'yyyy-MM-dd HH:mm:ss')}}</div>"}
            , {field: 'remark', width: 100, title: '备注', align: 'center'}
            , {field: 'status', width: 100, title: '状态', align: 'center'}
            , {fixed: 'right', width: 280, title: '功能操作', align: 'center', toolbar: '#toolBar'}

        ];

        //【渲染TABLE】
        func.tableIns(cols, "tableList");

        //【设置弹框】
        func.setWin("抽样信息");

    }
});
