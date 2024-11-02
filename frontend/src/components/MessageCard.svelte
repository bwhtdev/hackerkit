<script lang='ts'>
  const { message } = $props();
  const { id, text, username, createdAt } = message;
</script>

<div class="cursor-pointer mb-4 max-w-sm bg-white dark:bg-slate-800 border rounded-lg shadow-sm p-7 border-slate-200/60 dark:border-slate-400">
  <div class="block mb-3 inline-flex items-center space-x-3">
    <a
      href={`/profile/${username}`}
      class="text-lg font-bold leading-none tracking-tight text-slate-900 dark:text-slate-200">
      @{username}
    </a>
    <p class='text-xs text-slate-400 dark:text-slate-500'>{new Date(createdAt).toDateString()}</p>
  </div>
  <a
    href={`/${id}`}
    class="block pb-4 text-slate-500 dark:text-slate-400">
    {text}
  </a>
</div>
