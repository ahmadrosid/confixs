<script>
    import Button from "../components/ui/Button.svelte";
    import RetroOutline from "../components/ui/RetroOutline.svelte";
    import PlusIcon from 'lucide-svelte/icons/plus';
    import BaseLayout from "../components/BaseLayout.svelte";
    
    // Sample data for projects
    let projects = [
        { id: 1, name: "Project A", description: "Setup database schema", progress: 70 },
        { id: 2, name: "Project B", description: "Finish deployment", progress: 40 },
        { id: 3, name: "Project C", description: "UI/UX Design", progress: 90 }
    ];
</script>

<BaseLayout>
    <div class="bg-pattern absolute inset-x-0 top-0 -z-0 h-[280px] border-b border-gray-500"></div>

    <div class="container mx-auto max-w-6xl px-4 pb-12 relative">
        <div class="py-12 pb-4">
            <h1 class="text-4xl font-bold">Projects</h1>
        </div>

        <div class="grid grid-cols-1 sm:grid-cols-3 gap-6 py-8">
            {#each projects as project}
                <RetroOutline>
                    <div class="p-6">
                        <div class="text-xl font-bold">{project.name}</div>
                        <div class="text-sm text-gray-600">{project.description}</div>
                        <div class="text-sm text-gray-600">Progress: {project.progress}%</div>
                        <Button className="mt-4">View Details</Button>
                    </div>
                </RetroOutline>
            {/each}
        </div>

        <div class="py-8">
            <RetroOutline className="w-full">
                <div class="border-b border-gray-800 p-3">
                    <div class="flex items-center justify-between">
                        <h2 class="text-xl font-bold">
                            Actions
                        </h2>
                        <div>
                            <Button className="inline-flex items-center gap-2 pl-3"> <PlusIcon size="20"/> Add</Button>
                        </div>
                    </div>
                </div>
                <div class="p-3">
                    <ul class="space-y-2 py-4 list-inside">
                        <li>
                            <h3 class="font-semibold">Create a new project</h3>
                        </li>
                        <li>
                            <h3 class="font-semibold">Update existing projects</h3>
                        </li>
                        <li>
                            <h3 class="font-semibold">Delete completed projects</h3>
                        </li>
                    </ul>
                </div>
            </RetroOutline>
        </div>
    </div>
</BaseLayout>