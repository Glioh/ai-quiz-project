<script lang="ts">
    import Leaderboard from "../../lib/Leaderboard.svelte";
    import { leaderboard } from "../../service/player/player";
    import Button from "../../lib/Button.svelte";

    function goHome() {
        window.location.href = "/";
    }

    $: console.log('PlayerEndView leaderboard value:', $leaderboard);
</script>

<div class="flex justify-center bg-purple-500 min-h-screen w-full">
    <div class="mt-32">
        <h2 class="text-center text-white text-5xl font-bold">Game ended!</h2>
        <div class="flex flex-wrap gap-2 mt-10">
            <Leaderboard finish={true} leaderboard={$leaderboard} />
        </div>
        <div class="flex justify-center mt-8">
            <div class="[&>button]:bg-purple-600 [&>button]:hover:bg-purple-700">
                <Button on:click={goHome}>Back to Home</Button>
            </div>
        </div>
    </div>
</div>