<script lang="ts">
  import BaseLayout from "../components/BaseLayout.svelte";
  import RetroOutline from "../components/ui/RetroOutline.svelte";
  import { Link, navigate } from "svelte-routing";
  import PocketBase from "pocketbase";

  let pb = new PocketBase('http://localhost:8090');

  let name = '';
  let email = '';
  let password = '';
  let confirmPassword = '';
  let termsAccepted = false;
  let error = '';
  let showPassword = false;
  let showConfirmPassword = false;

  async function handleRegister(event: Event) {
    event.preventDefault();

    if (password !== confirmPassword) {
      error = "Passwords do not match";
      return;
    }

    if (!termsAccepted) {
      error = "You must accept the Terms of Service and Privacy Policy";
      return;
    }

    try {
      const data = {
        name,
        email,
        password,
        passwordConfirm: confirmPassword,
      };

      await pb.collection('users').create(data);
      
      // Optionally, you can automatically log in the user after registration
      await pb.collection('users').authWithPassword(email, password);
      
      console.log('Registration successful');
      navigate('/dashboard'); // Redirect to dashboard or home page after successful registration
    } catch (err: any) {
      console.error('Registration failed:', err);
      error = err.message || 'Registration failed. Please try again.';
    }
  }

  function togglePasswordVisibility(field: 'password' | 'confirmPassword') {
    if (field === 'password') {
      showPassword = !showPassword;
    } else if (field === 'confirmPassword') {
      showConfirmPassword = !showConfirmPassword;
    }
  }
</script>

<style>
  .password-input {
    -webkit-text-security: disc;
    font-family: text-security-disc;
  }
  .password-input.visible {
    -webkit-text-security: none;
    font-family: inherit;
  }
</style>

<BaseLayout>
<div class="bg-white flex items-center justify-center min-h-[90vh]">
  <RetroOutline className='max-w-md'>
    <div class="bg-white p-6">
      <h1 class="text-2xl font-bold mb-6">Register</h1>

      <p class="mb-4">
        Already have an account? <Link
          to="/login"
          class="text-blue-600 hover:underline">Login</Link
        >
      </p>

      <form on:submit={handleRegister}>
        <div class="mb-4">
          <label for="name" class="block mb-2">
            Full Name <span class="text-red-500">*</span>
          </label>
          <input
            type="text"
            id="name"
            bind:value={name}
            class="w-full px-3 py-2 border border-gray-300"
            required
          />
        </div>

        <div class="mb-4">
          <label for="email" class="block mb-2">
            Email <span class="text-red-500">*</span>
          </label>
          <input
            type="email"
            id="email"
            bind:value={email}
            class="w-full px-3 py-2 border border-gray-300"
            required
          />
        </div>

        <div class="mb-4">
          <label for="password" class="block mb-2">
            Password <span class="text-red-500">*</span>
          </label>
          <div class="relative">
            <input
              type="text"
              id="password"
              bind:value={password}
              class="w-full px-3 py-2 border border-gray-300 pr-10 password-input"
              class:visible={showPassword}
              required
            />
            <span
              class="absolute inset-y-0 right-0 flex items-center pr-3 cursor-pointer"
              on:click={() => togglePasswordVisibility('password')}
            >
              <svg
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                stroke-width="1.5"
                stroke="currentColor"
                class="w-5 h-5 text-gray-500"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  d="M2.036 12.322a1.012 1.012 0 010-.639C3.423 7.51 7.36 4.5 12 4.5c4.638 0 8.573 3.007 9.963 7.178.07.207.07.431 0 .639C20.577 16.49 16.64 19.5 12 19.5c-4.638 0-8.573-3.007-9.963-7.178z"
                />
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"
                />
              </svg>
            </span>
          </div>
        </div>

        <div class="mb-6">
          <label for="confirm-password" class="block mb-2">
            Confirm Password <span class="text-red-500">*</span>
          </label>
          <div class="relative">
            <input
              type="text"
              id="confirm-password"
              bind:value={confirmPassword}
              class="w-full px-3 py-2 border border-gray-300 pr-10 password-input"
              class:visible={showConfirmPassword}
              required
            />
            <span
              class="absolute inset-y-0 right-0 flex items-center pr-3 cursor-pointer"
              on:click={() => togglePasswordVisibility('confirmPassword')}
            >
              <svg
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                stroke-width="1.5"
                stroke="currentColor"
                class="w-5 h-5 text-gray-500"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  d="M2.036 12.322a1.012 1.012 0 010-.639C3.423 7.51 7.36 4.5 12 4.5c4.638 0 8.573 3.007 9.963 7.178.07.207.07.431 0 .639C20.577 16.49 16.64 19.5 12 19.5c-4.638 0-8.573-3.007-9.963-7.178z"
                />
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"
                />
              </svg>
            </span>
          </div>
        </div>

        <div class="mb-6">
          <label class="flex items-center">
            <input type="checkbox" class="mr-2" bind:checked={termsAccepted} required />
            <span class="text-sm">
              I agree to the <a href="#" class="text-blue-600 hover:underline"
                >Terms of Service</a
              >
              and
              <a href="#" class="text-blue-600 hover:underline"
                >Privacy Policy</a
              >
            </span>
          </label>
        </div>

        {#if error}
          <div class="mb-4 text-red-500">{error}</div>
        {/if}

        <button
          type="submit"
          class="w-full bg-black text-white py-2 hover:bg-gray-800"
        >
          Register
        </button>
      </form>
    </div>
  </RetroOutline>
</div>
</BaseLayout>