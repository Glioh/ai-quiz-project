<script lang="ts">
    import { createEventDispatcher } from "svelte";
    import Button from "../../lib/Button.svelte";
    import type { PlayerGame } from "../../service/player/player";

    const dispatch = createEventDispatcher();

    let code: string = "";
    let name: string = "";
    export let game: PlayerGame;
    let showJoinForm = false;

    function join(){
        dispatch("join");
        game.join(code, name);
    }

    function selectPlayer() {
        showJoinForm = true;
    }

    function selectHost() {
        window.location.href = "/#/host";
    }
</script>

<div class="bg-purple-400 min-h-screen w-full flex items-center justify-center">
    {#if !showJoinForm}
        <!-- Layer 1: Home Page -->
        <div class="text-center">
            <h1 class="text-white font-bold text-6xl mb-8">Welcome to Quiz!</h1>
            <p class="text-white text-xl mb-12">Choose how you want to participate</p>
            <div class="flex gap-4 justify-center">
                <Button on:click={selectHost}>Host a Game</Button>
                <Button on:click={selectPlayer}>Join a Game</Button>
            </div>
        </div>
    {:else}
        <!-- Layer 2: Join Game Form -->
        <div>
            <h2 class="text-white font-bold text-5xl text-center">Quiz</h2>
            <div class="flex flex-col gap-2 mt-10 items-center">
                <input bind:value={code} type="text" placeholder="Game code" class="p-2 rounded" />
                <input bind:value={name} type="text" placeholder="Name" class="p-2 rounded" />
                <Button on:click={join}>Join game</Button>
            </div>
        </div>
    {/if}
</div>