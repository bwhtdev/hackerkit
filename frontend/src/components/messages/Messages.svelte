<script lang="ts">
  import MessageCard from './MessageCard.svelte';

  import { onMount } from 'svelte';
  import { messageData, messageStatus, reloadMessageData } from './messageStore';

  onMount(() => {
     reloadMessageData();
  });
</script>

{#if $messageStatus == 'loading'}
  <p>Loading messages...</p>
{:else if $messageStatus == 'error' || $messageData.error}
  <p>Cannot load messages</p>
  <p>{$messageData.error}</p>
{:else}
  {#each $messageData as message}
    <MessageCard {message} />
  {/each}
{/if}
