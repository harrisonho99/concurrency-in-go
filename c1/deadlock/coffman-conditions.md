# *Coffman Conditions*

* **One of conditions true when deadlock occurs**

> **Mutual Exclusion**
> A concurrent process holds exclusive rights to a resource at any one time.

> **Wait For Condition**
> A concurrent process must simultaneously hold a resource and be waiting for an additional resource.

> **No Preemption**
> A resource held by a concurrent process can only be released by that process, so it fulfills this condition.

> **Circular Wait**
> A concurrent process (P1) must be waiting on a chain of other concurrent pro‐ cesses (P2), which are in turn waiting on it (P1), so it fulfills this final condition too.
