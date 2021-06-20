package patterns.behavioral.iterator.example1.models;

import patterns.behavioral.iterator.example1.Product;

import java.util.Iterator;

public class DevStoreIterator implements Iterator {
    Product[] catalog;
    int position = 0;

    public DevStoreIterator(Product[] catalog) {
        this.catalog = catalog;
    }

    @Override
    public boolean hasNext() {
        return position < catalog.length && catalog[position] != null;
    }

    @Override
    public Object next() {
        return catalog[position++];
    }

    @Override
    public void remove() {
        if (position <= 0) {
            throw new IllegalStateException("Can't remove item until you've done at least one next()");
        }
        if (catalog[position-1] == null) {
            for (int i = position-1; i < (catalog.length - 1); i++) {
                catalog[i] = catalog[i+1];
            }
            catalog[catalog.length - 1] = null;
        }
    }
}
