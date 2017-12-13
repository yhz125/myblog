<div class="container">
    {{range .blogs}}
        <h4><a href="/blog/detail?id={{.Id}}">{{.Title}}</a></h4>
        <p >{{.Content}}</p>
        <hr/>
    {{end}}
</div>