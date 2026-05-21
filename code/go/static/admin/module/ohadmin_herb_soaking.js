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
            , {field: 'id', width: 80, title: 'ID', align: 'center', sort: true, fixed: 'left'}
            , {field: 'hospital_name', width: 250, title: '医院名称', align: 'center'}
            , {field: 'patient_name', width: 100, title: '患者姓名', align: 'center'}
            , {field: 'prescription_number', width: 100, title: '处方号', align: 'center'}
            , {field: 'start_time', width: 180, title: '泡药开始时间', align: 'center', templet:"<div>{{layui.util.toDateString(d.start_time*1000, 'yyyy-MM-dd HH:mm:ss')}}</div>"}
            , {field: 'end_time', width: 180, title: '泡药结束时间', align: 'center', templet:"<div>{{layui.util.toDateString(d.end_time*1000, 'yyyy-MM-dd HH:mm:ss')}}</div>"}
            , {field: 'duration', width: 100, title: '泡药时长', align: 'center'}
            , {field: 'soaking_person', width: 100, title: '泡药员工姓名', align: 'center'}
            , {field: 'word_content', width: 100, title: '工作内容', align: 'center'}
            , {field: 'barcode', width: 100, title: '条形码', align: 'center'}
            , {field: 'status', width: 100, title: '状态', align: 'center'}

        ];

        //【渲染TABLE】
        func.tableIns(cols, "tableList");

        //【设置弹框】
        func.setWin("泡药信息");

    }
});
