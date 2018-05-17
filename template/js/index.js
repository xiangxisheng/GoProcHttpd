window.index=function(){
  var table_htoa = function (tableObj) {
    if (!tableObj.hasOwnProperty('Columns')) {
      alert('have no tableObj.Columns')
      return
    }
    if (!tableObj.hasOwnProperty('Rows')) {
      alert('have no tableObj.Rows')
      return
    }
    var rows = []
    for (var line in tableObj.Rows) {
      var row = {}
      for (var k in tableObj.Columns) {
        row[tableObj.Columns[k]] = tableObj.Rows[line][k]
      }
      rows.push(row)
    }
    return rows
  }
  var dictionary_name_to_title = function(name) {
    var obj = {};
    obj['firadio_uc'] = '认证中心';
    obj['firadio_ucenter'] = '用户中心';
    obj['firadio_yun'] = '云服务';
    obj['email'] = '邮箱';
    obj['session'] = '会话';
    obj['user'] = '用户';
    obj['ntuser'] = '域用户';
    obj['tenpay'] = '财付通';
    obj['date'] = '日期';
    obj['process'] = '耗点记录';
    obj['log'] = '日志';
    obj['list'] = '列表';
    obj['gt10w'] = '超10万点数'
    obj['cdn'] = '分布式节点'
    obj['balance'] = '余额'
    obj['qquin'] = 'QQ号码'
    if (obj.hasOwnProperty(name)) {
      return obj[name];
    }
  }
  new Vue({
    el: '#app',
    created: function() {
      this.update()
    },
    data: {
      config: {
        project: '',
        module: '',
        object: '',
        action: '',
        test: 0
      },
      table: {
        hashRows: [],
        project: [],
        test: 0
      },
      PMOA: {
        project: {},
        module: {},
        object: {},
        action: {}
      }
    },
    methods: {
      click1: function () {
        this.update()
      },
      update: function () {
        this.$http.get('/proc', {}).then(function (res) {
          // alert(res.data)
          this.table.hashRows = table_htoa(res.data.Table)
          this.procTable(this.table.hashRows)
        }, function (res) {
          alert(res.status)
        })
      },
      rand: function (min, max) {
        return Math.floor(Math.random() * (max - min + 1)) + min
      },
      rand_char: function (charstr) {
        return charstr.substr(this.rand(0, charstr.length - 1), 1)
      },
      rand_arr: function (arr) {
        return arr[this.rand(0, arr.length - 1)]
      },
      getName: function (len) {
        var ret = ''
        for (var i = 0; i < len; i++) {
          ret += this.rand_char('abcdefghijklmnopqrstuvwxyz')
        }
        ret += '@firadio.com'
        return ret
      },
      procTable: function (tableObj) {
        for (var key in tableObj) {
          const row = tableObj[key]
          console.log(row.SPECIFIC_NAME)
          if (!this.procRow(row)) {
            
            continue
          }
          
          var key = row.SPECIFIC_SCHEMA
          if (!this.PMOA.project.hasOwnProperty(key)) {
            const obj = {}
            obj.name = row.SPECIFIC_SCHEMA
            obj.title = dictionary_name_to_title(obj.name)
            this.PMOA.project[key] = []
            this.table.project.push(obj)
          }
          var key = row.SPECIFIC_SCHEMA + '_' + row.module
          if (!this.PMOA.module.hasOwnProperty(key)) {
            const obj = {}
            obj.name = row.module
            obj.title = dictionary_name_to_title(obj.name)
            this.PMOA.project[row.SPECIFIC_SCHEMA].push(obj)
            this.PMOA.module[key] = []
          }
          var key = row.SPECIFIC_SCHEMA + '_' + row.module + '_' + row.object
          if (!this.PMOA.object.hasOwnProperty(key)) {
            const obj = {}
            obj.name = row.object
            obj.title = dictionary_name_to_title(obj.name)
            this.PMOA.module[row.SPECIFIC_SCHEMA + '_' + row.module].push(obj)
            this.PMOA.object[key] = []
          }
          var key = row.SPECIFIC_SCHEMA + '_' + row.module + '_' + row.object + '_' + row.action
          if (!this.PMOA.action.hasOwnProperty(key)) {
            const obj = {}
            obj.name = row.action
            obj.title = dictionary_name_to_title(obj.name)
            this.PMOA.object[row.SPECIFIC_SCHEMA + '_' + row.module + '_' + row.object].push(obj)
            this.PMOA.action[key] = []
          }
        }
        console.log(this.PMOA)
      },
      procRow: function (row) {
        arr = row.SPECIFIC_NAME.split('_')
        if (arr.length <= 1) return false
        row.module = arr[1]
        if (arr.length <= 2) return false
        row.object = arr[2]
        if (arr.length <= 3) return false
        row.action = arr[3]
        return true
      }
    }
  })
}
