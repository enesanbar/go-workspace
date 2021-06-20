package patterns.behavioral.iterator.example1.models;

import patterns.behavioral.iterator.example1.Product;

import java.util.ArrayList;
import java.util.Iterator;

public class GeekyStoreIterator implements Iterator {
    ArrayList<Product> catalog;
    int position = 0;

    public GeekyStoreIterator(ArrayList<Product> catalog) {
        this.catalog = catalog;
    }

    @Override
    public boolean hasNext() {
        return position < catalog.size() && catalog.get(position) != null;
    }

    @Override
    public Object next() {
        return catalog.get(position++);
    }

}
