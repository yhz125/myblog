<div class="container">
    <form method="post" action="/blog">
      <div class="form-group">
        <label for="title">标题</label>
        <input type="hidden" name="Id" value="{{.blog.Id}}">
        <input type="text" class="form-control" name="Title" value="{{.blog.Title}}" id="Title" placeholder="标题">
      </div>
      <div class="form-group">
        <label for="content">内容</label>
        <textarea rows="10" cols="20" class="form-control" name="Content" id="Content" >{{.blog.Content}}</textarea>         
      </div>      
      <button type="submit" class="btn btn-default">提交</button>
    </form>
</div>