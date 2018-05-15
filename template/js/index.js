window.index=function(){
  new Vue({
    el: '#app',
    created: function() {
      this.$http.get('/proc', {}).then(function (res) {
        // alert(res.data)
      }, function (res) {
        alert(res.status)
      });
    },
    data: {
      config: {
        cnxing: '李,王,张',
        count: 20
      },
      rows: []
    },
    methods: {
      click1: function () {
        this.rows = []
        for (var i = 0; i < this.config.count; i++) {
          var row = {}
          row.email = this.getName(8)
          row.firstName = this.rand_arr(this.config.cnxing.split(','))
          row.lastName = this.rand_arr(this.config.cnxing.split(','))
          this.rows.push(row)
        }
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
        ret += '@firadio.com';
        return ret
      }
    }
  })
}