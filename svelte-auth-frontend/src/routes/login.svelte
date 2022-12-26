<script>
  import { goto } from "@sapper/app.mjs";

  let email='', password=''

  const submitForm = async () => {
     await fetch('http://localhost:8000/api/login', {
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        credentials: "include",
        body : JSON.stringify({
          email, 
          password
        })
     });

     await goto('/');
  }
</script>

<form on:submit|preventDefault={submitForm}>
    <h1 class="h3 mb-3 fw-normal">Please sign in</h1>

    <div class="form-floating">
      <input bind:value={email} type="email" class="form-control" id="floatingInput" placeholder="name@example.com">
      <label for="floatingInput">Email address</label>
    </div>
    <div class="form-floating">
      <input bind:value={password} type="password" class="form-control" id="floatingPassword" placeholder="Password">
      <label for="floatingPassword">Password</label>
    </div>

    <button class="w-100 btn btn-lg btn-primary" type="submit">Sign in</button>
</form>