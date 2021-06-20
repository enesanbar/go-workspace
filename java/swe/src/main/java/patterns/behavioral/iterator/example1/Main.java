package patterns.behavioral.iterator.example1;

import patterns.behavioral.iterator.example1.models.Catalog;
import patterns.behavioral.iterator.example1.models.Seller;

public class Main {

    public static void main(String[] args) {

        Catalog devStoreCatalog = new DevStoreCatalog();
        Catalog geekyStoreCatalog = new DevStoreCatalog();

        Seller seller = new Seller(geekyStoreCatalog, devStoreCatalog);
        seller.printCatalog();
    }

}
