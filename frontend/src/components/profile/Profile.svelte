<script lang='ts'> 
  let { username } = $props();

  import { loggedIn, username as currentUsername } from '@components/authStore';

  interface User {
    id: string;
    username: string;
    name: string;
    avatar: string;
    description: string;
    createdAt: string;
  }
  
  const getUser = async (): Promise<User> => {
    const res = await fetch(`/api/v1/user/username/${username}`, {
      headers: {'Access-Control-Allow-Origin': '*'}
    });
    return await res.json();
  };
</script>

{#await getUser()}
  <p>Loading user profile...</p>
{:then user}
  {#if !user.error}
    <div>
      <h1>{name}</h1>
      <p>{username} - {new Date(user.createdAt).toDateString()}</p>
      <img src={user.avatar} width='100' height='100' />
    </div>

    {#if $loggedIn && username == $currentUsername}
      <p>Hello, {username}!</p>
    {/if}
  {:else}
    <p>User does not exist.</p>
  {/if}
{:catch error}
  <p>Cannot load user profile.</p>
  <p>{error.message}</p>
{/await}  
