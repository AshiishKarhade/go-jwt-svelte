<script>
  import { goto } from "@sapper/app.mjs";

  let name='', email='', password=''

  const submitForm = async () => {
     await fetch('http://localhost:8000/api/register', {
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        body : JSON.stringify({
          name, 
          email, 
          password
        })
     });

     await goto('/login');
  }
</script>

<form on:submit|preventDefault={submitForm}>
    <h1 class="h3 mb-3 fw-normal">Register</h1>

    <div class="form-floating">
      <input bind:value={name} type="text" class="form-control" id="floatingInput" placeholder="John Doe" required>
      <label for="floatingInput">Full Name</label>
    </div>
    <div class="form-floating">
      <input bind:value={email} type="email" class="form-control" id="floatingInput" placeholder="name@example.com" required>
      <label for="floatingInput">Email address</label>
    </div>
    <div class="form-floating">
      <input bind:value={password} type="password" class="form-control" id="floatingPassword" placeholder="Password" required>
      <label for="floatingPassword">Password</label>
    </div>

    <button class="w-100 btn btn-lg btn-primary" type="submit">Sign up</button>
</form>