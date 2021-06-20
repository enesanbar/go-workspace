package solid.ocp.example;

public interface Specification<T> {

    boolean isSatisfied(T item);

}
