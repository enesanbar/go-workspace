package patterns.behavioral.chain.example1;

public class Main {

    public static void main(String[] args) {
        CEOPurchasePower ceoPurchasePower = new CEOPurchasePower();
        DirectorPurchasePower directorPurchasePower = new DirectorPurchasePower();
        ManagerPurchasePower managerPurchasePower = new ManagerPurchasePower();

        managerPurchasePower.setSuccessor(directorPurchasePower);
        directorPurchasePower.setSuccessor(ceoPurchasePower);

        PurchaseRequest purchaseRequest = new PurchaseRequest(20001, "new product");
        managerPurchasePower.processRequest(purchaseRequest);
    }

}
