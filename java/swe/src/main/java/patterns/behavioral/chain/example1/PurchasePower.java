package patterns.behavioral.chain.example1;

public abstract class PurchasePower {

    protected static final double BASE = 1000;
    protected PurchasePower successor;

    abstract protected double getAllowable();
    abstract protected String getRole();

    public PurchasePower getSuccessor() {
        return successor;
    }

    public void setSuccessor(PurchasePower successor) {
        this.successor = successor;
    }

    public void processRequest(PurchaseRequest request) {
        System.out.println("Processing purchase request: " + request.getAmount());

        if (request.getAmount() < this.getAllowable()) {
            System.out.println(this.getRole() + " will approve $" + request.getAmount());
        } else if (successor != null) {
            System.out.println(this.getRole() + " cannot approve prices larger than $" + this.getAllowable());
            System.out.println("Delegating the job to " + successor.getRole());
            successor.processRequest(request);
        }
    }
}
