# Network Aware Scheduling in Kubernetes

The Kuberenets Scheduler is part of the controlplane that helps schedule pods to appropriate nodes.
The Scheduler uses two processes to find the best node for a pod to be scheduled. The two steps are filtering and scoring. 
- Filtering : The scheduler finds a node that is meets the requirements requested by the user.
- Scoring : In order to find the best node, a scoring mechanism is used to select the best-fit

While the Kubernetes Scheduler takes CPU and memory into account, studies have shown that network resources are expensive resources that would should also be considered. In this project, we have configured the default scheduler to take network resources into account. 

```
    spec:
      hostNetwork: true
      schedulerName: my-scheduler
      containers:
      - name: img-python
        image: stevekim01310/kimjangbak:image
        imagePullPolicy: IfNotPresent
        stdin: true
        tty: true
        ports:
        - containerPort: 5000
      restartPolicy: Always
```
## Notes

- More than one node is required for the scheduler to work. 
