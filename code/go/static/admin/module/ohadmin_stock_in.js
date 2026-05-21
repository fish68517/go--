/**
 * 处方管理
 * @author mw
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
            , {field: 'id', width: 70, title: 'id', align: 'center', sort: true, fixed: 'left'}
            , {field: 'order_no', width: 150, title: '入库单号', align: 'center'}
            , {field: 'warehouse_id', width: 120, title: '仓库ID', align: 'center'}
            , {field: 'operator_id', width: 100, title: '操作人ID', align: 'center'}
            , {field: 'create_time', width: 70, title: '创建时间', align: 'center'}
            , {field: 'remark', width: 70, title: '备注', align: 'center'}
            , {fixed: 'right', width: 280, title: '功能操作', align: 'center', toolbar: '#toolBar'}
        ];

        //【渲染TABLE】
        func.tableIns(cols, "tableList");

        //【设置弹框】
        func.setWin("入库管理");

        //【设置状态】
        func.formSwitch('status', null, function (data, res) {
            console.log("开关回调成功");
        });
    }
});
