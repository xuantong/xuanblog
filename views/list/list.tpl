
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
          {{ .Contenthtml}}
        </p>
        <p class="text-right">
          <a class="btn btn-primary" href="/p/{{ .Id}}">查看更多 »</a>
        </p>
      </div>
{{end}}
      <div class="pagination pagination-centered">
        <ul>
          <li>
            <a href="#">上一页</a>
          </li>
          <li>
            <a href="#">1</a>
          </li>
          <li>
            <a href="#">2</a>
          </li>
          <li>
            <a href="#">3</a>
          </li>
          <li>
            <a href="#">4</a>
          </li>
          <li>
            <a href="#">5</a>
          </li>
          <li>
            <a href="#">下一页</a>
          </li>
        </ul>
      </div>