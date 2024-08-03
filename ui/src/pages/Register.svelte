<script lang="ts">
  import BaseLayout from "../components/BaseLayout.svelte";
  import RetroOutline from "../components/ui/RetroOutline.svelte";
  import { Link, navigate } from "svelte-routing";
  import PocketBase from "pocketbase";
  import EyeIcon from 'lucide-svelte/icons/eye'

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
            <button
              class="absolute inset-y-0 right-0 flex items-center pr-3 cursor-pointer"
              on:click={() => togglePasswordVisibility('password')}
              tabindex="0"
              type="button"
            >
              <EyeIcon className='size-4 text-gray-500'/>
            </button>
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
            <button
              class="absolute inset-y-0 right-0 flex items-center pr-3 cursor-pointer"
              on:click={() => togglePasswordVisibility('confirmPassword')}
              type="button"
            >
              
            <EyeIcon className='size-4 text-gray-500'/>
            </button>
          </div>
        </div>

        <div class="mb-6">
          <label class="flex items-center">
            <input type="checkbox" class="mr-2" bind:checked={termsAccepted} required />
            <span class="text-sm">
              I agree to the <a href="/" class="text-blue-600 hover:underline"
                >Terms of Service</a
              >
              and
              <a href="/" class="text-blue-600 hover:underline"
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