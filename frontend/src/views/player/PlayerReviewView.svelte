<script lang="ts">
    import { points } from "../../service/player/player";
    import AudioController from "../../lib/AudioController.svelte";
    import { onMount, tick } from 'svelte';

    let audioController: AudioController;
    $: correct = $points > 0;

    onMount(async () => {
        // Wait for next tick to ensure audioController is initialized
        await tick();
        
        try {
            if (audioController) {
                if (correct) {
                    audioController.playCorrect();
                } else {
                    audioController.playIncorrect();
                }
            }
        } catch (err) {
            console.log("Error playing sound:", err);
        }
    });
</script>

<div
    class="min-h-screen text-white w-full relative {correct
        ? 'bg-green-500'
        : 'bg-red-600'} flex justify-center items-center"
>
    <div class="absolute top-8 right-8">
        <AudioController 
            bind:this={audioController}
            audioFile=""
            iconColor="white"
            hoverBgColor={correct ? "bg-green-600" : "bg-red-700"}
        />
    </div>

    {#if correct}
        <div class="text-center">
            <h2 class="text-3xl font-bold">Correct!</h2>
            <p class="text-2xl">+ {$points} points</p>
        </div>
    {:else}
        <h2 class="text-3xl">Incorrect!</h2>
    {/if}
</div>