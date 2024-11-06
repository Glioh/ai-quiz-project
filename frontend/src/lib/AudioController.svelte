<script lang="ts">
    import { onMount, onDestroy } from 'svelte';
    import { faVolumeUp, faVolumeMute } from '@fortawesome/free-solid-svg-icons';
    import Fa from 'svelte-fa';
    import { isMuted } from '../stores/audioStore';

    export let audioFile: string;
    export let iconColor: string = "white";
    export let hoverBgColor: string = "bg-purple-600";
    
    let audio: HTMLAudioElement | null = null;
    let correctAudio: HTMLAudioElement | null = null;
    let incorrectAudio: HTMLAudioElement | null = null;

    // Function to determine if audio should loop based on filename
    function shouldLoop(filename: string): boolean {
        return filename === "Play.mp3"; // Only Play.mp3 will loop
    }

    onMount(() => {
        // Main audio
        if (!audio && audioFile) {
            audio = new Audio(audioFile);
            audio.loop = shouldLoop(audioFile);
            audio.volume = 0.5;
            audio.muted = $isMuted;
            audio.play().catch(error => console.log("Audio autoplay failed:", error));
        }

        // Initialize correct/incorrect audio
        correctAudio = new Audio("Correct.mp3");
        incorrectAudio = new Audio("Incorrect.mp3");
        correctAudio.volume = 0.5;
        incorrectAudio.volume = 0.5;
    });

    onDestroy(() => {
        if (audio) {
            audio.pause();
            audio.currentTime = 0;
            audio = null;
        }
    });

    function toggleMute() {
        $isMuted = !$isMuted;
        if (audio) {
            audio.muted = $isMuted;
        }
        if (correctAudio) correctAudio.muted = $isMuted;
        if (incorrectAudio) incorrectAudio.muted = $isMuted;
    }

    // Handle audio file changes
    $: if (audio && audioFile && audio.src !== new URL(audioFile, window.location.href).href) {
        audio.src = audioFile;
        audio.loop = shouldLoop(audioFile);
        audio.play().catch(error => console.log("Audio play failed:", error));
    }

    // Export functions to play correct/incorrect sounds
    export function playCorrect() {
        if (correctAudio) {
            correctAudio.currentTime = 0;
            correctAudio.muted = $isMuted;
            correctAudio.play().catch(error => console.log("Correct sound play failed:", error));
        }
    }

    export function playIncorrect() {
        if (incorrectAudio) {
            incorrectAudio.currentTime = 0;
            incorrectAudio.muted = $isMuted;
            incorrectAudio.play().catch(error => console.log("Incorrect sound play failed:", error));
        }
    }
</script>

<button 
    class="p-2 rounded-full transition-colors {`text-${iconColor}`} hover:{hoverBgColor}"
    on:click={toggleMute}
    aria-label={$isMuted ? "Unmute" : "Mute"}
>
    <Fa icon={$isMuted ? faVolumeMute : faVolumeUp} size="1.5x"/>
</button>