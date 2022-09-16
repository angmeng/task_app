<script>
  export let meta = { "total_results": 0, "per_page": 10, "current_page": 0, "last_page": 0 }
</script>

<div class="pagination">
  {#if meta.current_page > 1}
    <a href="/" style="padding: 1em;">Previous</a>
  {:else}
    <a href="/" style="padding: 1em;" class="isDisabled">Previous</a>
  {/if}
  
  {#if meta.current_page < meta.last_page}
    <a href="/" style="padding: 1em;">Next</a>
  {:else}
    <a href="/" style="padding: 1em;"  class="isDisabled">Next</a>
  {/if}
</div>
<div>
  <p>
    Total Count: {meta.total_results}<br>
    Total Page: {(meta.last_page - meta.current_page) + 1}
  </p>
</div>

<style>
  div.pagination { padding-top: 10px;}
  
  .isDisabled {
                color: currentColor;
                cursor: not-allowed;
                opacity: 0.5;
                text-decoration: none;
              }
</style>