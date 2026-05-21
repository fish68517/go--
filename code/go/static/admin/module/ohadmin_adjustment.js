/**
 * 友链管理
 * @author mwq
 * @since 2021/7/26
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
            , {field: 'word_date', width: 180, title: '开始调剂时间', align: 'center', templet:"<div>{{layui.util.toDateString(d.word_date*1000, 'yyyy-MM-dd HH:mm:ss')}}</div>"}
            , {field: 'end_date', width: 180, title: '结束调剂时间', align: 'center', templet:"<div>{{layui.util.toDateString(d.end_date*1000, 'yyyy-MM-dd HH:mm:ss')}}</div>"}
            , {field: 'word_person', width: 100, title: '调剂员工姓名', align: 'center'}
            , {field: 'word_content', width: 100, title: '工作内容', align: 'center'}
            , {field: 'barcode', width: 100, title: '条形码', align: 'center'}
            , {field: 'status', width: 100, title: '状态', align: 'center'}

        ];

        //【渲染TABLE】
        func.tableIns(cols, "tableList");

        //【设置弹框】
        func.setWin("调剂信息");

    }
});
