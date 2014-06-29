
  {{range .articleList}}
      <div>
        <div class="page-header">
          <h2>
            <div>{{ .Title}}<small>{{dateformat .Pubdate "2006-01-02"}}</small></div>
          </h2>
          </div>
          <p>
          <span class="label">文字标签</span> <span class="label">文字标签</span> <span class="label">文字标签</span> <span class="label">文字标签</span>
        </p>
        <p class="lead">
          {{str2html .Abstract}}
        </p>
        <p class="text-right">
          <a class="btn btn-primary" href="/p/{{ .Id}}">查看更多 »</a>
        </p>
      </div>
{{end}}

{{.count}}

{{if gt .paginator.PageNums 1}}
      <div class="pagination pagination-centered">

        <ul>
          {{if .paginator.HasPrev}}
          <li><a href="{{.paginator.PageLinkFirst}}">{{"paginator.first_page"}}</a></li>
          <li><a href="{{.paginator.PageLinkPrev}}">&lt;</a></li>
          {{end}}
          {{range $index, $page := .paginator.Pages}}
            <li {{if $.paginator.IsActive .}} class="active"{{end}} >
                <a href="{{$.paginator.PageLink $page}}">{{$page}}</a>
            </li>
        {{end}}

        {{if .paginator.HasNext}}
          <li><a href="{{.paginator.PageLinkNext}}">下一页</a></li>
          <li><a href="{{.paginator.PageLinkLast}}">最后一页</a></li>
        {{end}}
        </ul>
      </div>
{{end}}
