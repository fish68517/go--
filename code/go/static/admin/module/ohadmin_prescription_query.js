/**
 * 处方查询
 * @author fxh
 * @since 2025/02/03
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
            , {field: 'id', width: 70, title: 'ID', align: 'center', sort: true, fixed: 'left'}
            , {field: 'hospital_name', width: 150, title: '医院名称', align: 'center'}
            , {field: 'prescription_number', width: 120, title: '处方号', align: 'center'}
            , {field: 'patient_name', width: 100, title: '患者姓名', align: 'center'}
            , {field: 'dosage', width: 70, title: '贴数', align: 'center'}
            , {field: 'decoction_method', width: 70, title: '年龄', align: 'center'}
            , {field: 'decoction_scheme', width: 100, title: '电话', align: 'center'}
            , {field: 'do_person', width: 100, title: '接方人员', align: 'center'}
            , {field: 'do_time', width: 120, title: '接方时间', align: 'center', templet:"<div>{{layui.util.toDateString(d.do_time*1000, 'yyyy-MM-dd HH:mm:ss')}}</div>"}
            , {field: 'pres_aduit_reviewer', width: 100, title: '审方人员', align: 'center'}
            , {field: 'pres_aduit_time', width: 120, title: '审方时间', align: 'center', templet:"<div>{{layui.util.toDateString(d.pres_aduit_time*1000, 'yyyy-MM-dd HH:mm:ss')}}</div>"}
            , {field: 'adjustment_reviewer', width: 100, title: '调剂人员', align: 'center'}
            , {field: 'adjustment_time', width: 120, title: '调剂时间', align: 'center', templet:"<div>{{layui.util.toDateString(d.adjustment_time*1000, 'yyyy-MM-dd HH:mm:ss')}}</div>"}
            , {field: 'audit_reviewer', width: 100, title: '复核人员', align: 'center'}
            , {field: 'audit_datetime', width: 120, title: '复核时间', align: 'center', templet:"<div>{{layui.util.toDateString(d.audit_datetime*1000, 'yyyy-MM-dd HH:mm:ss')}}</div>"}
            , {field: 'soaking_person', width: 100, title: '泡药人员', align: 'center'}
            , {field: 'soak_start_time', width: 120, title: '泡药时间', align: 'center', templet:"<div>{{layui.util.toDateString(d.soak_start_time*1000, 'yyyy-MM-dd HH:mm:ss')}}</div>"}
            , {field: 'decoction_person', width: 100, title: '煎药人员', align: 'center'}
            , {field: 'decoction_start_time', width: 120, title: '煎药时间', align: 'center', templet:"<div>{{layui.util.toDateString(d.decoction_start_time*1000, 'yyyy-MM-dd HH:mm:ss')}}</div>"}
            , {field: 'packing_personnel', width: 100, title: '包装人员', align: 'center'}
            , {field: 'pack_start_time', width: 120, title: '包装时间', align: 'center', templet:"<div>{{layui.util.toDateString(d.pack_start_time*1000, 'yyyy-MM-dd HH:mm:ss')}}</div>"}
            , {field: 'delivery_personnel', width: 100, title: '发货人员', align: 'center'}
            , {field: 'delivery_time', width: 120, title: '发货时间', align: 'center', templet:"<div>{{layui.util.toDateString(d.delivery_time*1000, 'yyyy-MM-dd HH:mm:ss')}}</div>"}
            , {field: 'current_state', width: 100, title: '处方状态', align: 'center'}

        ];

        //【渲染TABLE】
        func.tableIns(cols, "tableList");

        //【设置弹框】
        func.setWin("处方查询");


    }
});
layui.use(function(){
    var laydate = layui.laydate;
    laydate.render({
        elem: '#start_time'
    });
    laydate.render({
        elem: '#end_time'
    });
});

