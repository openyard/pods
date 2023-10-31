### Lamport Clock

---

Use logical timestamps as a version for a value to allow ordering of values across servers

<sup>References:</sup>  
<sup>https://martinfowler.com/articles/patterns-of-distributed-systems/lamport-clock.html </sup>

---

![Two servers each responsible for specific keys](assets/figure1.png "Figure 1")  
<sup>Figure 1: Two servers each responsible for specific keys</sup>

![Two servers each responsible for specific keys](assets/figure2.png "Figure 2")  
<sup>Figure 2: Two servers each responsible for specific keys</sup>

![Different leader folower groups storing different key values](assets/figure3.png "Figure 3")  
<sup>Figure 3: Different leader folower groups storing different key values</sup>

![Partial Order](assets/figure4.png "Figure 4")  
<sup>Figure 4: Partial Order</sup>

![Single leader-follower group saving key values](assets/figure5.png "Figure 5")  
<sup>Figure 5: Single leader-follower group saving key values</sup>