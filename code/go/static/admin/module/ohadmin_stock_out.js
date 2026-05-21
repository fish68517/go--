/**
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
            , {field: 'id', width: 70, title: 'ID', align: 'center', sort: true, fixed: 'left'}
            , {field: 'order_no', width: 180, title: '出库单号', align: 'center'}
            , {field: 'warehouse_id', width: 120, title: '仓库ID', align: 'center'}
            , {field: 'customer_id', width: 120, title: '客户ID', align: 'center', templet: function(d){
                    return d.customer_id || '-';
                }}
            , {field: 'out_type', width: 120, title: '出库类型', align: 'center', templet: function(d){
                    var typeMap = {1:'销售出库', 2:'退货出库', 3:'调拨出库', 4:'其他出库'};
                    return typeMap[d.out_type] || '未知类型';
                }}
            , {field: 'operator_id', width: 120, title: '操作人ID', align: 'center'}
            , {field: 'create_time', width: 180, title: '创建时间', align: 'center'}
            , {field: 'update_time', width: 180, title: '更新时间', align: 'center'}
            , {field: 'remark', minWidth: 150, title: '备注', align: 'left'}
            , {fixed: 'right', width: 280, title: '功能操作', align: 'center', toolbar: '#toolBar'}
        ];

        //【渲染TABLE】
        func.tableIns(cols, "tableList");

        //【设置弹框】
        func.setWin("出库管理");

        //【监听出库类型选择，控制客户ID显示】
        form.on('select(out_type)', function(data){
            var outType = data.value;
            if (outType == 1) { // 销售出库显示客户ID
                $('.customer-id-group').show();
            } else {
                $('.customer-id-group').hide();
                $('input[name="customer_id"]').val(''); // 清空客户ID
            }
        });

        //【设置状态开关（若有状态字段可启用，当前表无状态字段，此处预留）】
        // func.formSwitch('status', null, function (data, res) {
        //     console.log("开关回调成功");
        // });
    }
});
