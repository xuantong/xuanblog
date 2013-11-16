
<br/>
<h2>
	<div>{{ .article.Title}}<small>{{dateformat .article.Pubdate "2006-01-02"}}</small></div>
</h2> 
<hr/>
<br/>
{{str2html .article.Content}}
<hr/>
<div class="text-right">
	<small>PubTime:{{dateformat .article.Pubdate "2006-01-02"}} &nbsp; UpdateTime:{{dateformat .article.Updated "2006-01-02"}}</small>
</div>
<!-- Duoshuo Comment BEGIN -->
<div class="ds-thread"></div>
<script type="text/javascript">
var duoshuoQuery = {short_name:"xuantong"};//require,replace your short_name
(function() {
				var ds = document.createElement('script');
				ds.type = 'text/javascript';ds.async = true;
				ds.src = 'http://static.duoshuo.com/embed.js';
				ds.charset = 'UTF-8';
				(document.getElementsByTagName('head')[0] 
				|| document.getElementsByTagName('body')[0]).appendChild(ds);
})();
</script>
<!-- Duoshuo Comment END -->	
