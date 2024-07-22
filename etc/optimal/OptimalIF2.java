package optimal;

import java.util.HashMap;
import java.util.Map;

public class OptimalIF2 {

    static final Map<String, CheckInput> lookup = new HashMap<>();

    enum CheckInput {
        ID("id"),
        NAME("name"),
        AGE("age") {
            @Override
            boolean check(String value) {
                if (value.matches("[0-9]+")) {
                    return true;
                }
                return false;
            }
        };

        CheckInput(String key) {
            // TODO Auto-generated constructor stub
            lookup.put(key, this);
        }

        boolean check(String value) {
            return value != null && !value.isEmpty();
        }

    };

    public static void main(String[] args) {
        User user = new User();
        // user.id = "1";
        // user.name = "John";
        // user.age = 30;
        // CheckInput.ID.check(user.id);

    }

}

class User {
    private String id;
    private String name;
    private int age;

    public User(String id, String name, int age) {
        this.id = id;
        this.name = name;
        this.age = age;
    }

    public User() {
    }

    public String getId() {
        return id;
    }

    public void setId(String id) {
        this.id = id;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public int getAge() {
        return age;
    }

    public void setAge(int age) {
        this.age = age;
    }

}