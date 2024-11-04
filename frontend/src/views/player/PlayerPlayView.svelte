<script lang="ts">
    import { onDestroy } from "svelte";
    import QuizChoiceCard from "../../lib/play/QuizChoiceCard.svelte";
    import { COLORS } from "../../model/quiz";
    import type { PlayerGame } from "../../service/player/player";

    export let game: PlayerGame;
    let answered = false;
    let isGreen = true;
    let interval: number;

    function onClick(i: number) {
        game.answer(i);
        answered = true;
    }

    // Reactive statement to start the interval when `answered` becomes true
    $: if (answered && !interval) {
        interval = setInterval(() => {
            isGreen = !isGreen;
        }, 1000);
    }

    onDestroy(() => {
        if (interval) {
            clearInterval(interval);
        }
    });
</script>

<div class="flex flex-wrap w-full min-h-screen">
    {#if !answered}
        {#each COLORS as color, i}
            <QuizChoiceCard {color}>
                <button class="h-full w-full" on:click={() => onClick(i)}>
                    X
                </button>
            </QuizChoiceCard>
        {/each}
    {:else}
        <div class="min-h-screen w-full flex items-center justify-center transition-colors duration-500 ease-in-out"
             class:bg-green-500={isGreen}
             class:bg-red-500={!isGreen}>
            <div class="text-center">
                <h1 class="text-4xl font-bold text-white mb-4">Lightning Fast!</h1>
                <p class="text-xl text-white">Waiting for other players...</p>
            </div>
        </div>
    {/if}
</div>
