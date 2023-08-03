<div class="loader">
    <div class="loader__outer">
        <div class="loader__inner" />
        <div class="loader__inner" />
        <div class="loader__inner" />
    </div>
</div>

<style lang="scss">
    @keyframes loader_after {
        0%, 2.5% {
            translate: calc(-1 * (var(--width) / 2));
        }
        97.5%, 100% {
            translate: calc((var(--width) / 2));
        }
    }
    
    .loader {
        --width: 6em;
        --height: .5em;
        --blur-size: 0.0625em;
        --animation-speed: 1.5;
        --outer-box-shadow-size-1: .25em;
        --outer-box-shadow-size-2: .125em;

        opacity: .125;

        width: 100%;
        height: 100%;
        
        display: grid;
        grid-template-columns: 1fr var(--width) 1fr;
        grid-template-rows: 1fr var(--height) 1fr;
    }

    .loader__outer {
        grid-column: 2;
        grid-row: 2;

        overflow: hidden;
        border-radius: calc(var(--height) / 2);

        height: 100%;
        width: 100%;
        position: relative;

        background: #FFF8;
        box-shadow: 0 0 var(--outer-box-shadow-size-1) var(--outer-box-shadow-size-1) #fff2, 0 0 var(--outer-box-shadow-size-2) var(--outer-box-shadow-size-2) #fff1;
        
        // &::before {
        //     position: absolute;
        //     top: 0;
        //     bottom: 0;
        //     left: 0;
        //     right: 0;

        //     display: block;
        //     content: "";

        //     width: var(--width);
        //     height: var(--height);

        //     // border-radius: .25em;
        // }
    }

    .loader__inner {
        position: absolute;
        left: 50%;
        transform: translateX(-50%);

        display: block;
        content: "";

        background: #fff;
        box-shadow: 0 0 var(--blur-size) var(--blur-size) #fff;
        animation: loader_after calc(2s / var(--animation-speed)) infinite alternate;

        &:nth-child(1) {
            animation-timing-function: cubic-bezier(0, 0.25, 1, 0.75);
        }

        &:nth-child(2) {
            animation-timing-function: cubic-bezier(0, 0.6, 1, 1);
        }

        &:nth-child(3) {
            animation-timing-function: cubic-bezier(0, 0, 1, 0.4);
        }

        

        height: var(--height);
        width: var(--height);
        border-radius: calc(var(--height) / 2);
    }

    
</style>