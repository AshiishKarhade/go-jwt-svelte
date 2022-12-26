<script>
    import {authenticated} from "../stores/auth";
    import {goto} from "@sapper/app.mjs";
    let auth = false;
    authenticated.subscribe(a => auth = a);
    const logout = async () => {
        await fetch('http://localhost:8000/api/logout', {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            credentials: 'include',
        })
        await goto('/login');
    }
</script>

<nav class="navbar navbar-expand-md navbar-dark bg-dark mb-4">
  <div class="container-fluid">
    <a class="navbar-brand" href="/">Home</a>
    
    <div>
	{#if auth}
      <ul class="navbar-nav me-auto mb-2 mb-md-0">
        <li class="nav-item">
          <a class="nav-link" href="/" on:click={logout}>Logout</a>
        </li>
      </ul>
	  {:else}
	   <ul class="navbar-nav me-auto mb-2 mb-md-0">
        <li class="nav-item">
          <a class="nav-link" href="/login">Login</a>
        </li>
		<li class="nav-item">
          <a class="nav-link" href="/register">Register</a>
        </li>
      </ul>
	  {/if}
    </div>
  </div>
</nav>