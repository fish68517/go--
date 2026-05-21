/**
 * 接方管理
 * @author mwq
 * @since 2025/01/08
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
            , {field: 'hospitalName', width: 150, title: '医院名称', align: 'center'}
            , {field: 'prescriptionNumber', width: 120, title: '处方号', align: 'center'}
            , {field: 'patientName', width: 100, title: '患者姓名', align: 'center'}
            , {field: 'patientSex', width: 70, title: '性别', align: 'center'}
            , {field: 'patientAge', width: 70, title: '年龄', align: 'center'}
            , {field: 'patientPhone', width: 100, title: '电话', align: 'center'}
            , {field: 'patientAddress', width: 100, title: '地址', align: 'center'}
            , {field: 'departmentName', width: 120, title: '科室', align: 'center'}
            , {field: 'inpatientArea', width: 120, title: '病区', align: 'center'}
            , {field: 'wardName', width: 120, title: '病房', align: 'center'}
            , {field: 'sickBed', width: 120, title: '病床', align: 'center'}
            , {field: 'diagnosisResult', width: 120, title: '诊断结果', align: 'center'}
            , {field: 'dosage', width: 70, title: '贴数', align: 'center'}
            , {field: 'drugPickupTime', width: 120, title: '取药时间', align: 'center', templet:"<div>{{layui.util.toDateString(d.DrugPickupTime*1000, 'yyyy-MM-dd HH:mm:ss')}}</div>"}
            , {field: 'decoctionScheme', width: 70, title: '煎药方案', align: 'center'}
            , {field: 'oneTimeDosage', width: 110, title: '一煎时间', align: 'center'}
            , {field: 'twoTimeDosage', width: 110, title: '二煎时间', align: 'center'}
            , {field: 'packageCount', width: 110, title: '包装量', align: 'center'}
            , {field: 'doTime', width: 120, title: '接方时间', align: 'center', templet:"<div>{{layui.util.toDateString(d.DoTime*1000, 'yyyy-MM-dd HH:mm:ss')}}</div>"}
            , {field: 'distributionCompany', width: 120, title: '收件地址', align: 'center'}
            , {field: 'distributionAddress', width: 120, title: '收件地址', align: 'center'}
            , {field: 'distributionPhone', width: 120, title: '联系电话', align: 'center'}
            , {field: 'distributionType', width: 70, title: '快递类型', align: 'center'}
            , {field: 'orderTime', width: 120, title: '订单时间', align: 'center', templet:"<div>{{layui.util.toDateString(d.OrderTime*1000, 'yyyy-MM-dd HH:mm:ss')}}</div>"}
            , {field: 'reviewer', width: 100, title: '审核人', align: 'center'}
            , {field: 'reviewTime', width: 120, title: '审核时间', align: 'center', templet:"<div>{{layui.util.toDateString(d.ReviewTime*1000, 'yyyy-MM-dd HH:mm:ss')}}</div>"}
            , {field: 'auditStatus', width: 100, title: '审核状态', align: 'center'}
            , {fixed: 'right', width: 150, title: '功能操作', align: 'center', toolbar: '#toolBar'}
        ];

        //【渲染TABLE】
        func.tableIns(cols, "tableList");

        //【设置弹框】
        func.setWin("接方信息");

        //【设置状态】
        func.formSwitch('status', null, function (data, res) {
            console.log("开关回调成功");
        });
    }
});
