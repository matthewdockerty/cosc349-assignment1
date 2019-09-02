db = db.getSiblingDB("recipesdb");
db.createCollection("recipes");

use recipesdb;

db.recipes.insert({
	name: "Sous Vide Whole Chicken",
    method: "1. Preheat water to 150°F using a sous vide precision cooker (I use Anova sous vide).<br><br>2. Prepare chicken by rinsing under cold water. Pat dry with paper towels and set it aside.<br><br>3. Mix the dry rub ingredients on a big plate. Now, prepare a large sous vide bag or a Ziploc bag by folding the top of the bag back over itself to form a hem. This will prevent chicken seasonings from getting on the edges of the bag. Set aside.<br><br>4. Place the chicken on the plate and rub the mixture all over it. Sprinkle the excess dry rub inside the chicken if there’s any left.<br><br>5. Slide the chicken into the prepared bag. Unfold the edge before closing the bag. Seal the bag using either a vacuum sealer or a hand pump.<br><br>6. Lower your bagged chicken into the preheated water bath, making sure the whole chicken is under the waterline. If using Ziploc bag, slowly lower your bagged chicken into your water bath, letting the pressure of the water press air out through the top of the bag. Once most of the air is out of the bag, carefully seal the bag just above the waterline. Cook for 6 hours.<br><br>7. Once the chicken is done, remove from the water bath and transfer it onto a plate. Gently pat with paper towels. Preserve the cooking liquid from the bag if you like for serving or for flavorful chicken soup or chicken stock later.<br><br>8. Heat a cast iron skillet over high heat. Melt butter and sear the whole chicken on all sides until the skin is golden brown and crispy, about 5 minutes. You can also cut up the chicken first before searing.<br><br>9. Serve with fresh cilantro. Enjoy!<br><br><a href=\"https://www.foodista.com/recipe/PVDD5X6R/sous-vide-whole-chicken\">Recipe by Sharon Chen</a><br><a href=\"https://creativecommons.org/licenses/by/3.0/\">CC BY 3.0</a>",
    ingredients: [
        "1 whole chicken (4-5 pounds)",
        "2 tablespoons unsalted butter",
        "Fresh cilantro for serving",
        "2 teaspoons Kosher salt",
        "2 teaspoons paprika",
        "1 teaspoon cayenne pepper",
        "1 teaspoon dried thyme",
        "2 teaspoons ground black pepper",
        "1/2 teaspoon garlic powder"
	],
	image: cat("/vagrant/data/recipe1.b64")
});

db.recipes.insert({
    name: "Strawberries Chantilly",
    method: "1. Place in a large bowl the strawberries and sprinkle with sugar. Stir and let it rest covered in the refrigerator until ready to eat.<br><br>2. Meanwhile in a large bowl or the bowl of a KitchenAid with the whisk attachment, add the heavy cream, powder sugar and vanilla. Whisk until soft peaks form.<br><br>3. In a large glass place couple of tablespoons of Chantilly (from step 2), a layer of strawberries, a layer of Chantilly.<br><br><a href=\"https://www.foodista.com/recipe/L3YWYKQ7/strawberries-chantilly\">Recipe by Giangi Townsend</a><br><a href=\"https://creativecommons.org/licenses/by/3.0/\">CC BY 3.0</a>",
    ingredients: [
        "1 cup heavy cream",
        "¼ to 1/3 cup powder sugar",
        "1 teaspoon vanilla extract",
        "1 large basket of fresh strawberries, cleaned cored and quartered",
        "1 tablespoon granulated sugar (optional)"
    ],
    image: cat("/vagrant/data/recipe2.b64")
});

db.recipes.insert({
    name: "Buffalo Cauliflower Tacos",
    method: "1. For the Buffalo Sauce: In a small sauce pan add the hot sauce butter, vinegar, Worchester sauce, pepper, garlic powder and paprika and place over low heat. Bring to a simmer, whisking occasionally until the butter has melted. Remove room the heat and set aside to cool.<br><br>2. For the Cauliflower: Preheat the oven to 425 degrees F.<br><br>3. In a large mixing bowl add the cauliflower florets, olive oil. Garlic powder and chili powder and gently toss to coat. Add the Buffalo sauce and using a large rubber spatula, gently mix to distribute the sauce evenly.<br><br>4. Place the seasoned cauliflower on a sheet pan and spread out so the florets are not overcrowded. Place in the oven and bake for about 20 minutes, rotating the pan halfway through.<br><br>5. While the cauliflower is baking, reheat the additional Buffalo sauce and reserve.<br><br>6. Remove from the oven and drizzling with some additional Buffalo sauce.<br><br>7. To assemble the tacos: heat the fresh or store-bought tortillas slightly, place some shredded lettuce in the tortilla, followed by a generous scoop of roasted cauliflower. Drizzle the with ranch dressing and serve.<br><br><a href=\"https://www.foodista.com/recipe/448LTD7J/buffalo-cauliflower-tacos\">Recipe by Simmer + Sauce</a><br><a href=\"https://creativecommons.org/licenses/by/3.0/\">CC BY 3.0</a>",
    ingredients: [
        "Buffalo Sauce:",
        "1 1/4 cups Louisiana Hot Sauce",
        "1 cup unsalted butter",
        "2 tablespoons white vinegar",
        "1 teaspoon Worchester sauce",
        "1/4-1/2 teaspoon cayenne pepper",
        "1/2 teaspoon garlic powder",
        "1/4 teaspoon paprika",
        "For the Cauliflower:",
        "1 cauliflower head, cut into small bite-sized florets",
        "2 tablespoons Extra virgin olive oil",
        "1 teaspoon garlic powder",
        "1 teaspoon chili powder",
        "1/2 cup homemade Buffalo sauce (see above)",
        "8-10 small corn tortillas, for serving",
        "1 1/2 cups shredded iceberg lettuce, for serving",
        "Ranch dressing, for drizzling"
    ],
    image: cat("/vagrant/data/recipe3.b64")
});

db.recipes.insert({
    name: "Halibut in Soy, Honey, Ginger Marinate",
    method: "1. Make the marinate by adding the honey, soy sauce, rice vinegar, garlic, parsley ginger and dried chili peppers flakes in a food processor or blender until all well blended.<br><br>2. Season with salt and pepper both sides of the halibut pieces. Place them in a dish and pour the marinate, reserving a couple of tablespoons. Marinate up to 1 hour turning once.<br><br>3. Sear at medium high for 4 to 6 minutes per side or until it has a nice brown color. Drizzle remaining marinade over the fish.<br><br><a href=\"https://www.foodista.com/recipe/RDRVZN7W/halibut-in-soy-honey-ginger-marinate\">Recipe by Giangi Townsend</a><br><a href=\"https://creativecommons.org/licenses/by/3.0/\">CC BY 3.0</a>",
    ingredients: [
        "1/3 cup honey",
        "1/4 cup soy sauce",
        "1/4 cup seasoned rice vinegar",
        "1 garlic clove, minced",
        "1 two-inch fresh ginger, outer skin removed",
        "1/4 teaspoon dried red chili pepper flakes",
        "1/2 bunch fresh parsley, finely chopped",
        "6 halibut pieces"
    ],
    image: cat("/vagrant/data/recipe4.b64")
});

db.recipes.insert({
    name: "Chicken and Cilantro Cream Sauce",
    method: "1. Season the chicken tenders lightly with the salt, pepper, and garlic salt.<br><br>2. Preheat the vegetable oil in a large skillet over medium-high heat. Place the chicken in the pan and brown on both sides.<br><br>3. Remove from the pan, and add in the shallots and reduce the heat to medium. Cook the shallots, stirring frequently, until they start to soften up. About 2-3 minutes.<br><br>4. Add the chicken broth, lime juice, tequila, cilantro, red pepper flakes, heavy cream, and salt. Simmer over medium to medium-low heat for 2 minutes. Add chicken back into pan, and continue on a low simmer, and spoon the sauce over the chicken as it cooks. Cook an additional 3-4 minutes, or until the chicken is cooked through and no longer pink.<br><br><a href=\"https://www.foodista.com/recipe/D7WSFKYB/chicken-and-cilantro-cream-sauce\">Recipe by Nicole Johnson</a><br><a href=\"https://creativecommons.org/licenses/by/3.0/\">CC BY 3.0</a>",
    ingredients: [
        "1 pound chicken tenders boneless, skinless, and trimmed of any excess fat",
        "salt, pepper, and garlic salt lightly applied to chicken breasts",
        "2-3 tablespoons vegetable oil",
        "1/4 cup finely diced shallot",
        "1 cup chicken broth",
        "2 tablespoons fresh lime juice",
        "1 tablespoon good tequila",
        "1 1/2 tablespoons Gourmet Garden Lightly Dried cilantro OR Cilantro stir-in paste",
        "1/2 teaspoon Gourmet Garden Lightly Dried Red Pepper flakes",
        "1/4 cup heavy cream",
        "1/2 teaspoon salt"
    ],
    image: cat("/vagrant/data/recipe5.b64")
});

db.recipes.insert({
    name: "Bulgogi Pork",
    method: "1. Trim the pork, slice and place in a medium size mixing bowl. Set aside.<br><br>2. In a second small mixing bowl add the soy sauce, sesame oil, chili paste, brown sugar, garlic, salt, and pepper. Mix well to blend.<br><br>3. Add the marinate to the sliced pork and toss to coat. Place the pork and marinade in a large ziplock bag and marinate for at least 8 hours or overnight.<br><br>4. To finish the bulgogi, place a large cast iron skillet over high heat. When hot, but not smoking, a place a layer of pork in the skillet without overcrowding. Sear the pork until you begin to see nice coloration/caramelization, about 2-3 minutes. Using tongs, carefully turn the pork over and cook for an additional 2-3 minutes. Transfer the bulgogi to a plate and cover to keep warm. Repeat the above process with the remaining pork, discarding the remaining marinade.<br><br>5. To serve, place in a large serving bowl (with or without rice), sprinkle with toasted sesame seeds and scallions.<br><br><a href=\"https://www.foodista.com/recipe/Q3B4TMR2/bulgogi-pork\">Recipe by Simmer + Sauce</a><br><a href=\"https://creativecommons.org/licenses/by/3.0/\">CC BY 3.0</a>",
    ingredients: [
        "1 pounds pork shoulder, trimmed and cut into thin 2x1-inch slices",
        "1/3 cup plus 2 tablespoons soy sauce",
        "3 tablespoons toasted sesame oil",
        "1 teaspoon garlic chili paste",
        "2 tablespoons brown sugar",
        "3 cloves garlic, grated",
        "1/4 teaspoon salt",
        "1/4 teaspoon black pepper",
        "1 bunch green scallions, white and green parts thinly sliced",
        "1 tablespoon toasted white sesame seeds"
    ],
    image: cat("/vagrant/data/recipe6.b64")
});