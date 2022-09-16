<script>
  export let meta = { "total_results": 0, "per_page": 10, "current_page": 0, "last_page": 0 }
</script>

<div>
  {#if meta.current_page > 0}
    <a href="/" style="padding: 1em;">Previous</a>
  {/if}
  {#if meta.current_page < meta.last_page}
    <a href="/" style="padding: 1em;">Next</a>
  {/if}
</div>
<div>
  <p>
    Total Count: {meta.total_results}<br>
    Total Page: {(meta.last_page - meta.current_page) + 1}
  </p>
</div>