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
            {type: 'checkbox', fixed: 'left'},
            {field: 'id', width: 80, title: 'ID', align: 'center', sort: true, fixed: 'left'},
            {field: 'order_no', width: 180, title: '盘点单号', align: 'center'},
            {
                field: 'warehouse_id', width: 120, title: '仓库', align: 'center',

            },
            {
                field: 'operator_id', width: 120, title: '操作人', align: 'center',

            },
            {
                field: 'create_time', width: 170, title: '创建时间', align: 'center',

            },
            {
                field: 'update_time', width: 170, title: '更新时间', align: 'center',
               
            },
            {field: 'remark', minWidth: 200, title: '备注', align: 'center'},
            {fixed: 'right', width: 200, title: '功能操作', align: 'center', toolbar: '#toolBar'}
        ];

        //【渲染TABLE】
        func.tableIns(cols, "tableList");

        //【设置弹框】
        func.setWin("盘点管理");

    }
});

