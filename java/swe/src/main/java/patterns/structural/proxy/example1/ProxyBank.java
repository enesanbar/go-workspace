package patterns.structural.proxy.example1;

import java.util.ArrayList;
import java.util.List;

public class ProxyBank implements Bank {

    private RealBank realBank = new RealBank();
    private static List<String> bannedClients;

    static {
        bannedClients = new ArrayList<>();
        bannedClients.add("james");
        bannedClients.add("giulietta");
    }

    @Override
    public void withdraw(String clientName) throws Exception {
        if (bannedClients.contains(clientName.toLowerCase())) {
            throw new Exception("Access Denied");
        }

        realBank.withdraw(clientName);
    }

}
