<script lang="ts">
    // 1. Precious Metals (Ounces)
    let gold = $state(2350);
    let silver = $state(28);
    let plat = $state(950);
    let pall = $state(1100);

    // 2. Industrial Metals (MT/kg)
    let copper = $state(12000);
    let iron = $state(99);
    let titan = $state(45);

    // 3. Automated Totals (Svelte 5 Magic)
    let totalPrecious = $derived(gold + silver + plat + pall);
    let totalIndustrial = $derived(copper + iron + titan);
    let grandTotal = $derived(totalPrecious + totalIndustrial);

    // 4. Fetch the Go Logic
    let apiResult = $state({ gs_ratio: '0.00', gold_copper: '0' });

    async function calculate() {
        const res = await fetch(`/api/calculate?gold=${gold}&silver=${silver}&plat=${plat}&pall=${pall}&cu=${copper}&fe=${iron}&ti=${titan}`);
        apiResult = await res.json();
    }
</script>

<main class="min-h-screen flex items-center justify-center bg-[#050505] font-sans selection:bg-yellow-500/30 p-4">
  <div class="relative group w-full max-w-2xl">
    
    <!-- ДИЗАЙНЕРСКИЙ ГЛОУ (Золотой градиент) -->
    <div class="absolute -inset-1 bg-gradient-to-tr from-yellow-600 to-orange-600 rounded-[2.5rem] blur opacity-10 group-hover:opacity-25 transition duration-1000"></div>
    
    <!-- ОСНОВНАЯ КАРТОЧКА -->
    <div class="relative bg-black/40 border border-white/10 backdrop-blur-3xl p-10 rounded-[2.5rem] shadow-2xl">
      
      <!-- HEADER (В стиле Station 01) -->
      <header class="flex justify-between items-start mb-10">
        <div>
          <h1 class="text-white/30 text-[10px] uppercase tracking-[0.4em] font-bold mb-1">Metal Flex by Anna Shulik</h1>
          <p class="text-white/90 text-sm font-light tracking-widest uppercase">Metal-Flex Index</p>
        </div>
        <!-- Пульсирующая точка (Золотая) -->
        <div class="w-2 h-2 rounded-full bg-yellow-500 shadow-[0_0_15px_rgba(234,179,8,0.5)] animate-pulse"></div>
      </header>

      <!-- ТИТУЛЬНИК -->
      <h2 class="text-4xl font-extralight tracking-tighter text-white mb-10 select-none">
        Forge <span class="text-yellow-500">Precision</span>
      </h2>

      <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
    <!-- PRECIOUS METALS (Золото, Серебро, Платина, Палладий) -->
    <section class="space-y-6">
        <h2 class="text-[10px] font-bold tracking-[0.3em] text-yellow-500/50 uppercase mb-4">Precious Assets (troy oz)</h2>
        
        <div class="space-y-4">
            <div class="relative">
                <span class="absolute -top-2 left-3 bg-[#050505] px-1 text-[8px] text-white/40 uppercase tracking-widest">Gold (XAU)</span>
                <input type="number" bind:value={gold} class="w-full bg-white/5 border border-white/5 p-3 rounded-xl text-white outline-none focus:border-yellow-500/50 transition-colors" />
            </div>

            <div class="relative">
                <span class="absolute -top-2 left-3 bg-[#050505] px-1 text-[8px] text-white/40 uppercase tracking-widest">Silver (XAG)</span>
                <input type="number" bind:value={silver} class="w-full bg-white/5 border border-white/5 p-3 rounded-xl text-white outline-none focus:border-white/20 transition-colors" />
            </div>

            <div class="relative">
                <span class="absolute -top-2 left-3 bg-[#050505] px-1 text-[8px] text-white/40 uppercase tracking-widest">Platinum (XPT)</span>
                <input type="number" bind:value={plat} class="w-full bg-white/5 border border-white/5 p-3 rounded-xl text-white outline-none focus:border-white/20 transition-colors" />
            </div>

            <div class="relative">
                <span class="absolute -top-2 left-3 bg-[#050505] px-1 text-[8px] text-white/40 uppercase tracking-widest">Palladium (XPD)</span>
                <input type="number" bind:value={pall} class="w-full bg-white/5 border border-white/5 p-3 rounded-xl text-white outline-none focus:border-white/20 transition-colors" />
            </div>
        </div>
    </section>

    <!-- INDUSTRIAL METALS (Медь, Железо, Титан) -->
    <section class="space-y-6">
        <h2 class="text-[10px] font-bold tracking-[0.3em] text-orange-500/50 uppercase mb-4">Industrial Bulk</h2>
        
        <div class="space-y-4">
            <div class="relative">
                <span class="absolute -top-2 left-3 bg-[#050505] px-1 text-[8px] text-white/40 uppercase tracking-widest">Copper (MT)</span>
                <input type="number" bind:value={copper} class="w-full bg-white/5 border border-white/5 p-3 rounded-xl text-white outline-none focus:border-orange-500/50 transition-colors" />
            </div>

            <div class="relative">
                <span class="absolute -top-2 left-3 bg-[#050505] px-1 text-[8px] text-white/40 uppercase tracking-widest">Iron Ore (MT)</span>
                <input type="number" bind:value={iron} class="w-full bg-white/5 border border-white/5 p-3 rounded-xl text-white outline-none focus:border-orange-500/50 transition-colors" />
            </div>

            <div class="relative">
                <span class="absolute -top-2 left-3 bg-[#050505] px-1 text-[8px] text-white/40 uppercase tracking-widest">Titanium (kg)</span>
                <input type="number" bind:value={titan} class="w-full bg-white/5 border border-white/5 p-3 rounded-xl text-white outline-none focus:border-orange-500/50 transition-colors" />
            </div>

            <button onclick={calculate} class="w-full bg-white/90 hover:bg-yellow-500 text-black font-bold py-4 rounded-2xl transition-all active:scale-95 uppercase text-xs tracking-widest mt-6 shadow-[0_0_20px_rgba(255,255,255,0.1)]">
                Sync Go-Engine
            </button>
        </div>
    </section>
</div>

      <!-- РЕЗУЛЬТАТЫ (Precision Results) -->
      <div class="mt-10 pt-8 border-t border-white/5 grid grid-cols-2 gap-6">
        <div class="relative group/stat">
          <p class="text-[10px] text-white/20 uppercase tracking-widest mb-2">G-S Ratio</p>
          <p class="text-3xl font-light text-white/90">{apiResult.gs_ratio}</p>
        </div>
        <div class="relative group/stat">
          <p class="text-[10px] text-white/20 uppercase tracking-widest mb-2">Grand Index</p>
          <p class="text-3xl font-light text-yellow-500">${grandTotal.toLocaleString()}</p>
        </div>
      </div>

      <!-- FOOTER (Твой фирменный стиль) -->
      <footer class="mt-12 flex flex-col items-center gap-3 pt-8 border-t border-white/5">
        <div class="flex items-center gap-1.5 group/love cursor-default">
          <span class="text-[10px] text-white/20 uppercase tracking-[0.2em] font-medium">Made with</span>
          <span class="text-rose-500 animate-pulse text-xs">❤️</span>
          <span class="text-[10px] text-white/20 uppercase tracking-[0.2em] font-medium">for you</span>
        </div>
        <div class="flex items-center gap-3 opacity-10">
          <span class="text-[8px] font-mono text-yellow-500">Go v1.22</span>
          <div class="w-1 h-1 rounded-full bg-white/20"></div>
          <span class="text-[8px] font-mono text-orange-400">Svelte 5</span>
        </div>
      </footer>

    </div>
  </div>
</main>
