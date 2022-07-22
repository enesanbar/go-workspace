from functools import wraps

from flask import Blueprint, render_template, jsonify, request, redirect, url_for, flash, get_flashed_messages

from catalog.application.cache.catalog import CatalogCache
from catalog.application.repositories.catalog import CatalogRepository
from catalog.domain.requests.create_category import CreateCategoryRequest
from catalog.domain.requests.create_product import CreateProductRequest
from catalog.domain.requests.get_product import GetProductsRequest


catalog = Blueprint('catalog', __name__)


def template_or_json(template=None):
    """"Return a dict from your view and this will either
    pass it to a template or render json. Use like:

    @template_or_json('template.html')
    """
    def decorated(f):
        @wraps(f)
        def decorated_fn(*args, **kwargs):
            ctx = f(*args, **kwargs)
            if request.headers.get("X-Requested-With") == "XMLHttpRequest" or not template:
                return jsonify(ctx)
            else:
                return render_template(template, **ctx)
        return decorated_fn
    return decorated


@catalog.context_processor
def product_name_processor():
    def full_name(product):
        return '{0} / {1}'.format(product.category.name, product.name)

    def flash_messages(**kwargs):
        return get_flashed_messages(**kwargs)

    return {'full_name': full_name, 'flash_messages': flash_messages}


# blueprint level filter
@catalog.app_template_filter('full_name')
def full_name_filter(product):
    return '{0} / {1}'.format(product['category'], product['name'])


@catalog.route('/')
@catalog.route('/home')
@template_or_json('home.html')
def home(catalog_repo: CatalogRepository):
    page = int(request.args.get('page', 0))
    per_page = int(request.args.get('per_page', 15))
    result = catalog_repo.get_products(GetProductsRequest(page, per_page))

    return {'count': len(result)}


@catalog.get('/products/<id>')
def product(id: int, catalog_repo: CatalogRepository, cache: CatalogCache):
    p = catalog_repo.get_product_by_id(id)
    product_key = f"product-{p.id}"
    cache.set(product_key, p.name, expiration=600)
    return render_template('product.html', product=p)


@catalog.route('/recent-products')
def recent_products(cache: CatalogCache):
    keys = cache.keys('product-*')
    recent = [cache.get(k).decode('utf-8') for k in keys]
    return jsonify({'products': recent})


@catalog.get('/products')
def products(catalog_repo: CatalogRepository):
    page = int(request.args.get('page', 0))
    per_page = int(request.args.get('per_page', 15))
    result = catalog_repo.get_products(GetProductsRequest(page, per_page))
    return render_template('products.html', products=result)


@catalog.post('/products')
def create_product(catalog_repo: CatalogRepository):
    name = request.form.get('name')
    price = float(request.form.get('price'))
    cat_name = request.form.get('category')
    cat = catalog_repo.get_category_by_name(cat_name)
    catalog_repo.create_product(CreateProductRequest(name, price, cat))
    flash(f'The product {name} has been created', 'success')
    # todo: return inserted ID
    return redirect(url_for('catalog.product', id=1))


@catalog.get('/product-create')
def render_create_product():
    return render_template('product-create.html')


@catalog.post('/category-create')
def create_category(catalog_repo: CatalogRepository):
    name = request.form.get('name')
    catalog_repo.create_category(CreateCategoryRequest(name))
    return "category created"


@catalog.get('/categories')
def categories(catalog_repo: CatalogRepository):
    result = catalog_repo.get_categories()
    return render_template('categories.html', categories=result)
    # return Response(
    #     json.dumps(get_categories, cls=CategoryJsonEncoder),
    #     mimetype="application/json",
    #     status=200,
    # )


@catalog.route('/category/<name>')
def category(name: str, catalog_repo: CatalogRepository):
    result = catalog_repo.get_category_by_name(name)
    return render_template('category.html', category=result)