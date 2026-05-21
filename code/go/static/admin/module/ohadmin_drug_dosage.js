/**
 * 药品超剂量配置管理
 * @author xxx
 * @since 2025/10/28
 */
layui.use(['func'], function () {

    // 声明变量
    var func = layui.func
        , form = layui.form
        , $ = layui.$;

    if (A == 'index') {
        // 【TABLE列数组】
        var cols = [
            {type: 'checkbox', fixed: 'left'}
            , {field: 'id', width: 100, title: '记录ID', align: 'center'}
            , {field: 'drug_id', width: 150, title: '药品ID', align: 'center'}
            , {field: 'max_dosage', width: 120, title: '最大剂量', align: 'center'}
            , {
                field: 'is_enabled',
                width: 100,
                title: '是否启用',
                align: 'center',
                templet: function(d) {
                    return d.is_enabled == 1 ? '<span class="layui-badge layui-bg-green">启用</span>' : '<span class="layui-badge layui-bg-gray">禁用</span>';
                }
            }
            , {field: 'dosage_desc', width: 200, title: '剂量说明', align: 'center'}
            , {field: 'create_time', width: 180, title: '创建时间', align: 'center'}
            , {field: 'update_time', width: 180, title: '更新时间', align: 'center'}
            , {fixed: 'right', width: 220, title: '功能操作', align: 'center', toolbar: '#toolBar'}
        ];

        // 【渲染TABLE】
        func.tableIns(cols, "tableList");

        // 【设置弹框】
        func.setWin("药品超剂量配置");
    }
});
