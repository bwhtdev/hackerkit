<script lang='ts'>
  import SignUpBtn from './SignUpBtn.svelte';
  import LogInBtn from './LogInBtn.svelte';
  import CreateMessageBtn from './CreateMessageBtn.svelte';

  import { onMount } from 'svelte';
  import { get } from 'svelte/store';

  import { username, loggedIn } from '@components/authStore';
  
  import { isLoggedIn, logOut, getUsername } from '../alpine/authUtils';
  
  onMount(() => {
    loggedIn.set(isLoggedIn());
    if (get(loggedIn)) username.set(getUsername());
  });
</script>

{#if $loggedIn}
  <CreateMessageBtn />
  <p class='mx-2'>{$username}</p>
  <button on:click={logOut} class="inline-flex items-center justify-center px-4 py-2 text-base font-medium leading-6 text-slate-600 dark:text-slate-200 whitespace-no-wrap bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-500 rounded-md shadow-sm hover:bg-slate-50 dark:hover:bg-slate-700 focus:outline-none focus:shadow-none">Log Out</button>
{:else}
  <LogInBtn />
  <SignUpBtn />
{/if}
