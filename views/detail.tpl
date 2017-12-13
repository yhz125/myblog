<div class="container">
    <div class="panel panel-default">
      <div class="panel-heading">
        <h3 class="panel-title">{{.detail.Title}}    <a class="btn btn-default" href="/blog/edit?id={{.detail.Id}}" role="button">编辑</a></h3>
      </div>
      <div class="panel-body">
        {{.detail.Content}}
      </div>
    </div>
</div>