/**
 * 库存管理
 * @author fengxh
 * @since 2025/10/30
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
            , {field: 'id', width: 180, title: '序号', align: 'center'}
            , {field: 'product_id', width: 120, title: '药品ID', align: 'center'}
            , {field: 'product_name', width: 150, title: '药品名称', align: 'center'}
            , {field: 'warehouse_id', width: 120, title: '库房ID', align: 'center'}
            , {field: 'quantity', width: 120, title: '库存数量', align: 'center'}
            , {field: 'created_time', width: 200, title: '创建时间', align: 'center'}
            , {field: 'updated_time', width: 200, title: '更新时间', align: 'center'}
            , {field: 'remark', width: 150, title: '备注', align: 'center'}
            , {fixed: 'right', width: 280, title: '功能操作', align: 'center', toolbar: '#toolBar'}

        ];

        //【渲染TABLE】
        func.tableIns(cols, "tableList");

        //【设置弹框】
        func.setWin("库存信息");

    }
});
