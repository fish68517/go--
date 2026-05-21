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
            , {field: 'id', width: 180, title: '序号', align: 'center'}
            , {field: 'equipment_type', width: 90, title: '设备类型', align: 'center'}
            , {field: 'device_name', width: 90, title: '设备名称', align: 'center'}
            , {field: 'device_room', width: 100, title: '煎药室', align: 'center'}
            , {field: 'unit_number', width: 120, title: '机组编号', align: 'center'}
            , {field: 'remark', width: 100, title: '备注', align: 'center'}
            , {fixed: 'right', width: 280, title: '功能操作', align: 'center', toolbar: '#toolBar'}

        ];

        //【渲染TABLE】
        func.tableIns(cols, "tableList");

        //【设置弹框】
        func.setWin("设备信息");

    }
});
