//値をグラフに表示させる
Chart.plugins.register({
  afterDatasetsDraw: function (chart, easing) {
      var ctx = chart.ctx;

      chart.data.datasets.forEach(function (dataset, i) {
          var meta = chart.getDatasetMeta(i);
          if (!meta.hidden) {
              meta.data.forEach(function (element, index) {
                  // 値の表示
                  ctx.fillStyle = 'rgb(255, 255, 255)';//文字の色
                  var fontSize = 12;//フォントサイズ
                  var fontStyle = 'normal';//フォントスタイル
                  var fontFamily = 'Arial';//フォントファミリー
                  ctx.font = Chart.helpers.fontString(fontSize, fontStyle, fontFamily);

                  var dataString = dataset.data[index].toString();
        
                  // 値の位置
                  ctx.textAlign = 'center';//テキストを中央寄せ
                  ctx.textBaseline = 'middle';//テキストベースラインの位置を中央揃え

                  var padding = 5;//余白
                  var position = element.tooltipPosition();
                  ctx.fillText(dataString, position.x, position.y - (fontSize / 2) - padding);
  
              });
          }
      });
  }
});


//=========== 円グラフ ============//
$('#chart01').on('inview', function(event, isInView) {//画面上に入ったらグラフを描画
if (isInView) {
  const ratio = document.getElementById("ratio")
  console.log(ratio.textContent)
  console.log(typeof(ratio.textContent))

  const remaining = ratio.textContent
  const lived = 100 - remaining


var ctx=document.getElementById("chart01");//グラフを描画したい場所のid
var chart=new Chart(ctx,{
type:'pie',//グラフのタイプ
data:{//グラフのデータ
labels:["生きた時間","残りの時間"],//データの名前
datasets:[{
    label:"生きた時間と残りの時間の比率",//グラフのタイトル
    backgroundColor:["#ff0000","#0000ff"],//グラフの背景色
    data:[lived.toFixed(2),remaining]//データ
  }]
},

options:{//グラフのオプション
maintainAspectRatio: false,//CSSで大きさを調整するため、自動縮小をさせない
legend:{
  display:true//グラフの説明を表示
},
tooltips:{//グラフへカーソルを合わせた際の詳細表示の設定
  callbacks:{
      label: function (tooltipItem, data) {
    return data.labels[tooltipItem.index]+ ": "+ data.datasets[0].data[tooltipItem.index] + "%";//%を最後につける
  }
  },    
},
title:{//上部タイトル表示の設定
  display: true,
  fontSize:10,
  text: '単位：%'
},
}
});

}
});