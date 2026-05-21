/*
 * 设备巡检记录管理
 * @author mwq
 * @since 2025/10/28
 */
layui.use(['func'], function () {

    //声明变量
    var func = layui.func
        , form = layui.form
        , $ = layui.$;

    if (A == 'index') {
        //【TABLE列数组】
        var cols = [
            {type: 'checkbox', fixed: 'left'}
            , {field: 'id', width: 80, title: 'ID', align: 'center', sort: true, fixed: 'left'}
            , {field: 'equipment_type', width: 150, title: '设备类型', align: 'center'}
            , {field: 'equipment_id', width: 150, title: '设备编号', align: 'center'}
            , {field: 'health_status', width: 120, title: '卫生状态', align: 'center', templet: function (d) {
                    var status = d.health_status;
                    if (status == 1) return '<span class="layui-badge layui-bg-green">良好</span>';
                    if (status == 0) return '<span class="layui-badge layui-bg-red">较差</span>';
                    return '<span class="layui-badge">未知</span>';
                }
            }
            , {field: 'disinfection_status', width: 120, title: '消毒状态', align: 'center', templet: function (d) {
                    var status = d.disinfection_status;
                    if (status == 1) return '<span class="layui-badge layui-bg-green">已消毒</span>';
                    if (status == 0) return '<span class="layui-badge layui-bg-red">未消毒</span>';
                    return '<span class="layui-badge">未知</span>';
                }
            }
            , {field: 'status', width: 120, title: '运行状态', align: 'center', templet: function (d) {
                    var status = d.status;
                    if (status == 1) return '<span class="layui-badge layui-bg-green">正常</span>';
                    if (status == 0) return '<span class="layui-badge layui-bg-red">异常</span>';
                    return '<span class="layui-badge">未知</span>';
                }
            }
            , {field: 'inspection_time', width: 180, title: '巡检时间', align: 'center', templet: "<div>{{layui.util.toDateString(d.inspection_time*1000, 'yyyy-MM-dd HH:mm:ss')}}</div>"}
            , {field: 'inspector', width: 120, title: '巡检人员', align: 'center'}
            , {fixed: 'right', width: 150, title: '功能操作', align: 'center', toolbar: '#toolBar'}
        ];

        //【渲染TABLE】
        func.tableIns(cols, "tableList");

        //【设置弹框】
        func.setWin("设备巡检记录");

        //【设置状态】
        func.formSwitch('status', null, function (data, res) {
            console.log("运行状态开关回调成功");
        });
    }
});
