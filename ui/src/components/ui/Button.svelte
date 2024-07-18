<script lang="ts">
    type Variant = 'primary' | 'secondary' | 'warning' | 'danger' | 'success' | 'outline-primary' | 'outline-secondary' | 'outline-danger' | 'outline-success';
    type Size = 'sm' | 'base' | 'md' | 'lg';

    export let className = "";
    export let variant: Variant = "primary";
    export let size: Size = "base";

    const baseClasses = "inline-flex items-center justify-center font-medium hover:underline";
    
    const variantClasses: Record<Variant, string> = {
        primary: "bg-black text-white hover:bg-gray-900",
        secondary: "bg-gray-200 text-black hover:bg-gray-300",
        danger: "bg-rose-500 text-white hover:bg-rose-600",
        warning: "bg-orange-500 text-white hover:bg-orange-600",
        success: "bg-green-500 text-white hover:bg-green-600",
        'outline-primary': "bg-transparent text-black border border-black",
        'outline-secondary': "bg-transparent text-gray-600 border border-gray-300 hover:bg-gray-100",
        'outline-danger': "bg-transparent text-rose-500 border border-rose-500 hover:bg-rose-50",
        'outline-success': "bg-transparent text-green-500 border border-green-500 hover:bg-green-50"
    };

    const sizeClasses: Record<Size, string> = {
        sm: "text-xs px-2 py-1",
        base: "text-sm px-3 py-1.5",
        md: "text-sm px-4 py-2",
        lg: "text-base px-6 py-3"
    };

    $: buttonClasses = `${baseClasses} ${variantClasses[variant]} ${sizeClasses[size]} ${className}`;
</script>

<button class={buttonClasses} on:click {...$$props}>
    <slot></slot>
</button>