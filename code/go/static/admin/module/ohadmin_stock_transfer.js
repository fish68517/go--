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
            , {field: 'order_no', width: 180, title: '调拨单号', align: 'center', sort: true}
            , {field: 'from_warehouse_id', width: 130, title: '调出仓库ID', align: 'center', sort: true}
            , {field: 'to_warehouse_id', width: 130, title: '调入仓库ID', align: 'center', sort: true}
            , {field: 'total_quantity', width: 120, title: '总数量', align: 'center', sort: true}
            , {field: 'operator_id', width: 110, title: '操作人ID', align: 'center'}
            , {field: 'create_time', width: 170, title: '创建时间', align: 'center', sort: true}
            , {field: 'update_time', width: 170, title: '更新时间', align: 'center', sort: true}
            , {field: 'remark', minWidth: 150, title: '备注', align: 'left'}
            , {fixed: 'right', width: 220, title: '功能操作', align: 'center', toolbar: '#toolBar'}
        ];

        //【渲染TABLE】
        func.tableIns(cols, "tableList");

        //【设置弹框】
        func.setWin("调拨管理");

        //【设置状态】
        func.formSwitch('status', null, function (data, res) {
            console.log("开关回调成功");
        });
    }
});
